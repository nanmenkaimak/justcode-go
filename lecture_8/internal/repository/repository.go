package repository

import "github.com/nanmenkaimak/justcode-go/lecture_8/internal/models"

type DatabaseRepo interface {
	CreateUser(newUser models.User) (int, error)
	Authenticate(email string, password string) (int, error)
	UpdateUser(user models.User) (bool, error)
	DeleteUserByID(id int) (bool, error)
}
