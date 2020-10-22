package user

import (
	"github.com/hrz8/go-take-arest/models"
	"gorm.io/gorm"
)

// Repository is an interface of User domain for user model method
type (
	Repository interface {
		GetAll(db *gorm.DB) (*[]models.User, error)
	}
)
