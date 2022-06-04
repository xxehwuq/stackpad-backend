package service

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type Service struct {
	Notebook NotebookService
}

func New(storage *storage.Storage) *Service {
	return &Service{
		Notebook: newNotebookService(storage.Notebook),
	}
}
