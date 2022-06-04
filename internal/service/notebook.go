package service

import (
	"fmt"

	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type NotebookService interface {
	Test()
}

type notebookService struct {
	storage storage.NotebookStorage
}

func newNotebookService(storage storage.NotebookStorage) *notebookService {
	return &notebookService{
		storage: storage,
	}
}

func (s *notebookService) Test() {
	fmt.Println("Test() from notebook service")
	s.storage.Test()
}
