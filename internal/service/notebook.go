package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type NotebookService interface {
	Add(notebook entity.Notebook, userId string) (string, error)
}

type notebookService struct {
	storage storage.NotebookStorage
}

func newNotebookService(storage storage.NotebookStorage) *notebookService {
	return &notebookService{
		storage: storage,
	}
}

func (s *notebookService) Add(notebook entity.Notebook, userId string) (string, error) {
	id, _ := gonanoid.New()

	notebook.Id = id

	err := s.storage.Add(notebook)
	if err != nil {
		return "", err
	}

	return id, nil
}
