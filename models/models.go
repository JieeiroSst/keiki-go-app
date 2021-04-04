package models

type User struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Book struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price uint `json:"price"`
}