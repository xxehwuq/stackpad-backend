package service

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
)

type Service struct {
	User UserService
}

type Pkg struct {
	PasswordManager hash.PasswordManager
}

func New(storage *storage.Storage, pkg Pkg) *Service {
	return &Service{
		User: newUserService(storage.User, pkg.PasswordManager),
	}
}
