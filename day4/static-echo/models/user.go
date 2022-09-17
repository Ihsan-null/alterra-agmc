package models

type User struct {
	Id       uint   `json:"user_id" param:"user_id"`
	Name     string `json:"name" param:"name"`
	Email    string `json:"email" param:"email" validate:"required,email"`
	Password string `json:"password" param:"password" validate:"required,min=8"`
}