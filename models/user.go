package models

import (
	"time"
)

// User will return given object about user
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	FullName  string    `gorm:"size:255;not null;" json:"full_name"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Address   *string   `gorm:"size:150;" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
