package service

import (
	"net/http"

	"github.com/hrz8/go-take-arest/domains/user"
	"github.com/hrz8/go-take-arest/shared/lib"
	"github.com/labstack/echo/v4"
)

type (
	handler struct {
		repository user.Repository
	}
)

// InitUserService will return User REST
func InitUserService(e *echo.Echo, repo user.Repository) {
	h := handler{
		repository: repo,
	}

	e.GET("/api/v1/users", h.GetAll)
}

func (h handler) GetAll(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	users, err := h.repository.GetAll(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  "db not connected",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   users,
	})
}
