package service

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/jwt"
)

type Service struct {
	User UserService
}

type Pkg struct {
	PasswordManager hash.PasswordManager
	JwtManager      jwt.JwtManager
}

func New(storage *storage.Storage, pkg Pkg) *Service {
	return &Service{
		User: newUserService(storage.User, pkg.PasswordManager, pkg.JwtManager),
	}
}
