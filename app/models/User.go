package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name		string	`json:"name"`
	Password	string	`json:"-"`
	Email		string	`json:"email"`
}
