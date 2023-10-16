package handlers

import (
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/repository"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/repository/dbrepo"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Repo *Repository

type Repository struct {
	DB      repository.DatabaseRepo
	DBRedis *redis.Client
}

func NewRepo(db *gorm.DB, dbRedis *redis.Client) *Repository {
	return &Repository{
		DB:      dbrepo.NewPostgresRepo(db),
		DBRedis: dbRedis,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
