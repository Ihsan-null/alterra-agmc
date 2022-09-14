package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name string `json:"name" query:"name"`
	Email string `json:"email" query:"email"`
	Password string `json:"password" query:"password"`
}