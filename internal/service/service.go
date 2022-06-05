package service

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type Service struct {
	User UserService
}

func New(storage *storage.Storage) *Service {
	return &Service{
		User: newUserService(storage.User),
	}
}
