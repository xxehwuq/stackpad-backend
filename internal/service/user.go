package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
)

type UserService interface {
	SignUp(user entity.User) (string, error)
}

type userService struct {
	storage         storage.UserStorage
	passwordManager hash.PasswordManager
}

func newUserService(storage storage.UserStorage, passwordManager hash.PasswordManager) *userService {
	return &userService{
		storage:         storage,
		passwordManager: passwordManager,
	}
}

func (s *userService) SignUp(user entity.User) (string, error) {
	id, _ := gonanoid.New()
	hashedPassword := s.passwordManager.Hash(user.Password)

	user.Id = id
	user.Password = hashedPassword

	err := s.storage.Add(user)
	if err != nil {
		return "", err
	}

	token := id // ! implement generating jwt token

	return token, nil
}
