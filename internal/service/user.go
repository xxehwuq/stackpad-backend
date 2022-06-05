package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type UserService interface {
	SignUp(user entity.User) (string, error)
}

type userService struct {
	storage storage.UserStorage
}

func newUserService(storage storage.UserStorage) *userService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) SignUp(user entity.User) (string, error) {
	id, _ := gonanoid.New()
	user.Id = id

	err := s.storage.Add(user)
	if err != nil {
		return "", err
	}

	token := id // ! implement generating jwt token

	return token, nil
}
