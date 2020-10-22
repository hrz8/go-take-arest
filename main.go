package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hrz8/go-take-arest/shared/config"
	"github.com/hrz8/go-take-arest/shared/container"
	"github.com/hrz8/go-take-arest/shared/database"
	"github.com/hrz8/go-take-arest/shared/lib"

	UserRepository "github.com/hrz8/go-take-arest/domains/user/repository"
	UserService "github.com/hrz8/go-take-arest/domains/user/service"
)

func main() {
	e := echo.New()

	appContainer := container.DefaultContainer()
	appConfig := appContainer.MustGet("shared.config").(config.AppConfigInterface)
	mysql := appContainer.MustGet("shared.mysql").(database.MysqlInterface)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	mysqlSess, err := mysql.Connect()
	if err != nil {
		panic(fmt.Sprintf("[ERROR] failed open mysql connection: %s", err.Error()))
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &lib.AppContext{
				Context:      c,
				MysqlSession: mysqlSess,
			}
			return next(ac)
		}
	})

	userRepository := UserRepository.NewUserRepository(mysqlSess)
	UserService.InitUserService(e, userRepository)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.GetAppPort())))
}
