package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/repository"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
