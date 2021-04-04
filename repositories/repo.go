package repositories

import (
	"github.com/JIeeiroSst/go-app/models"
)

type UserRepository interface {
	Login(username,password string) error
	Singup(username,password string) error
}

type BookRepository interface {
	FindAll() ([]models.Book,error)
	FindById(id int) (models.Book,error)
	Create(book models.Book) error
	Update(id int,book models.Book) error
	Delete(id int) error
}

type Repository interface {
	UserRepository
	BookRepository
}