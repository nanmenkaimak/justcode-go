package dbrepo

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
	"gorm.io/gorm"
	"time"
)

func (m *gistRepository) CreateGist(newGist models.GistRequest) (uuid.UUID, error) {
	err := m.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&newGist.Gist)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		newGist.Commit.GistID = newGist.Gist.ID
		res = tx.Create(&newGist.Commit)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		for i := 0; i < len(newGist.Files); i++ {
			newGist.Files[i].CommitID = newGist.Commit.ID
		}

		res = tx.Create(&newGist.Files)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return newGist.Gist.ID, nil
}

func (m *gistRepository) GetAllGists() ([]models.GistRequest, error) {
	var allGistsReq []models.GistRequest

	var allGists []models.Gist

	err := m.DB.Find(&allGists).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allGists); i++ {
		var lastCommit models.Commit
		err = m.DB.Where("gist_id = ?", allGists[i].ID).Order("created_at desc").Find(&lastCommit).Limit(1).Error
		if err != nil {
			return nil, err
		}
		var allFiles []models.File
		err = m.DB.Where("commit_id = ?", lastCommit.ID).Find(&allFiles).Error
		if err != nil {
			return nil, err
		}

		res := models.GistRequest{
			Gist:   allGists[i],
			Commit: lastCommit,
			Files:  allFiles,
		}
		allGistsReq = append(allGistsReq, res)
	}

	return allGistsReq, nil
}

func (m *gistRepository) GetGistByID(id uuid.UUID) (models.GistRequest, error) {
	var gistReq models.GistRequest

	var gist models.Gist

	err := m.DB.Where("id = ?", id).Find(&gist).Error
	if err != nil {
		return gistReq, err
	}

	var lastCommit models.Commit
	err = m.DB.Where("gist_id = ?", gist.ID).Order("created_at desc").Find(&lastCommit).Limit(1).Error
	if err != nil {
		return gistReq, err
	}
	var allFiles []models.File
	err = m.DB.Where("commit_id = ?", lastCommit.ID).Find(&allFiles).Error
	if err != nil {
		return gistReq, err
	}

	gistReq = models.GistRequest{
		Gist:   gist,
		Commit: lastCommit,
		Files:  allFiles,
	}

	return gistReq, nil
}

func (m *gistRepository) UpdateGist(updatedGist models.GistRequest) error {
	err := m.DB.Transaction(func(tx *gorm.DB) error {
		updatedGist.Gist.UpdatedAt = time.Now()
		res := tx.Model(&updatedGist.Gist).Where("id = ?", updatedGist.Gist.ID).Updates(updatedGist.Gist)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		updatedGist.Commit.GistID = updatedGist.Gist.ID
		res = tx.Create(&updatedGist.Commit)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		for i := 0; i < len(updatedGist.Files); i++ {
			updatedGist.Files[i].CommitID = updatedGist.Commit.ID
		}

		res = tx.Create(&updatedGist.Files)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		return nil
	})
	return err
}

func (m *gistRepository) DeleteGistByID(id uuid.UUID) error {
	err := m.DB.Where("id = ?", id).Delete(&models.Gist{}).Error
	return err
}
