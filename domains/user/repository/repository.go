package repository

import (
	"github.com/hrz8/go-take-arest/domains/user"
	"github.com/hrz8/go-take-arest/models"
	"gorm.io/gorm"
)

type (
	handler struct {
		db *gorm.DB
	}
)

// NewUserRepository return implementation of methods in transaction.Repositoru
func NewUserRepository(db *gorm.DB) user.Repository {
	return &handler{
		db: db,
	}
}

func (h handler) GetAll(db *gorm.DB) (*[]models.User, error) {
	var err error
	users := []models.User{}
	err = db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}
	return &users, err
}
