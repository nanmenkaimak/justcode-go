package dbrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/repository"
)

type postgresDBRepo struct {
	DB *sqlx.DB
}

func NewPostgresRepo(conn *sqlx.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
