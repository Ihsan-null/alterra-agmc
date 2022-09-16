package models

type Book struct {
	Id     int    `query:"id" validate:"required"`
	Title  string `query:"title" validate:"required"`
	Author string `query:"author" validate:"required"`
	UserId uint
}

type UpdateBook struct {
	Title  string `json:"title" query:"title" validate:"required"`
	Author string `json:"author" query:"author" validate:"required"`
}
