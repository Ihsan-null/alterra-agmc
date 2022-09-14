package models

type Book struct {
	Id int `json:"id" query:"id"`
	Title string `json:"title" query:"title"`
	Author string `json:"author" query:"author"`
}

type UpdateBook struct {
	Title string `json:"title" query:"title"`
	Author string `json:"author" query:"author"`
}