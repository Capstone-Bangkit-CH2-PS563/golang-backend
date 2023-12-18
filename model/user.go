package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password,omitempty" form:"password"`
	Token    string `gorm:"-"`
	Role     string `json:"role" form:"role" gorm:"type:enum('USER', 'ADMIN');default:'USER'"`
}
