package storage

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type NotebookStorage interface {
	Add(notebook entity.Notebook) error
}

type notebookStorage struct {
	db *gorm.DB
}

func newNotebookStorage(db *gorm.DB) *notebookStorage {
	return &notebookStorage{
		db: db,
	}
}

func (s *notebookStorage) Add(notebook entity.Notebook) error {
	result := s.db.Create(&notebook)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
