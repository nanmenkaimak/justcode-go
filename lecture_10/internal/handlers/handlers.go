package handlers

import (
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/repository"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/repository/dbrepo"
	"gorm.io/gorm"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
