package storage

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type NotebookStorage interface {
	Add(notebook entity.Notebook) error
	Get(userId string) ([]entity.Notebook, error)
	GetById(notebookId, userId string) (entity.Notebook, error)
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

func (s *notebookStorage) Get(userId string) ([]entity.Notebook, error) {
	var notebooks []entity.Notebook

	result := s.db.Where("user_id = ?", userId).Find(&notebooks)
	if result.Error != nil {
		return nil, result.Error
	}

	return notebooks, nil
}

func (s *notebookStorage) GetById(notebookId, userId string) (entity.Notebook, error) {
	var notebook entity.Notebook

	result := s.db.Take(&notebook, "id = ? AND user_id = ?", notebookId, userId)
	if result.Error != nil {
		return entity.Notebook{}, result.Error
	}

	return notebook, nil
}
