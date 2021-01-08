package users

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Password  string         `json:"-"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
