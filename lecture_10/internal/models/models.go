package models

import (
	"github.com/google/uuid"
	"time"
)

type GistRequest struct {
	Gist   Gist   `json:"gist"`
	Commit Commit `json:"commit"`
	Files  []File `json:"files"`
}

type Gist struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:varchar(50); unique"`
	Description string    `json:"description" gorm:"type:varchar(150)"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:now()"`
}

type Commit struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid; default:gen_random_uuid()"`
	GistID    uuid.UUID `json:"gist_id"`
	Gist      Gist      `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Comment   string    `json:"comment" gorm:"type:varchar(250)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

type File struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid; default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(50)"`
	CommitID  uuid.UUID `json:"commit_id"`
	Commit    Commit    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}
