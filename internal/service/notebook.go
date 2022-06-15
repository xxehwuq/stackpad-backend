package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type NotebookService interface {
	Add(notebook entity.Notebook, userId string) (string, error)
	GetAll(userId string) ([]entity.Notebook, error)
	GetById(notebookId, userId string) (entity.Notebook, error)
	DeleteById(notebook entity.Notebook) error
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
	notebook.UserId = userId

	err := s.storage.Add(notebook)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *notebookService) GetAll(userId string) ([]entity.Notebook, error) {
	notebooks, err := s.storage.GetAll(userId)
	if err != nil {
		return nil, err
	}

	return notebooks, nil
}

func (s *notebookService) GetById(notebookId, userId string) (entity.Notebook, error) {
	notebook, err := s.storage.GetById(notebookId, userId)
	if err != nil {
		return entity.Notebook{}, err
	}

	return notebook, nil
}

func (s *notebookService) DeleteById(notebook entity.Notebook) error {
	err := s.storage.DeleteById(notebook)
	if err != nil {
		return err
	}

	return nil
}
