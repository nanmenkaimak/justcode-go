package dbrepo

import (
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/repository"
	"gorm.io/gorm"
)

type gistRepository struct {
	DB *gorm.DB
}

func NewPostgresRepo(conn *gorm.DB) repository.Repository {
	return &gistRepository{
		DB: conn,
	}
}
