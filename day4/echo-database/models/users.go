package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestUpdateUser struct {
	Name     string `json:"name" query:"name"`
	Email    string `json:"email" query:"email" validate:"required,email"`
	Password string `json:"password" query:"password" validate:"required,min=8"`
}

type Login struct {
	Email    string `json:"email" query:"email" validate:"required,email"`
	Password string `json:"password" query:"password" validate:"required,min=8"`
}
