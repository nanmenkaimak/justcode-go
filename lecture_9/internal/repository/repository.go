package repository

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
)

type DatabaseRepo interface {
	CreateGist(newGist models.GistRequest) (uuid.UUID, error)
	GetAllGists() ([]models.GistRequest, error)
	GetGistByID(id uuid.UUID) (models.GistRequest, error)
	UpdateGist(updatedGist models.GistRequest) error
	DeleteGistByID(id uuid.UUID) error
}
