package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/jwt"
)

type UserService interface {
	SignUp(user entity.User) (string, error)
	SignIn(email, password string) (string, error)
	Confirm(userId string) error
}

type userService struct {
	storage         storage.UserStorage
	passwordManager hash.PasswordManager
	jwtManager      jwt.JwtManager
}

func newUserService(storage storage.UserStorage, passwordManager hash.PasswordManager, jwtManager jwt.JwtManager) *userService {
	return &userService{
		storage:         storage,
		passwordManager: passwordManager,
		jwtManager:      jwtManager,
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

	token, err := s.jwtManager.GenerateToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) SignIn(email, password string) (string, error) {
	password = s.passwordManager.Hash(password)

	user, err := s.storage.GetByCredentials(email, password)
	if err != nil {
		return "", err
	}

	token, err := s.jwtManager.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Confirm(userId string) error {
	err := s.storage.Confirm(userId)
	if err != nil {
		return err
	}

	return nil
}
