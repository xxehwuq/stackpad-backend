package storage

import (
	"errors"

	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type UserStorage interface {
	Add(user entity.User) error
	GetByCredentials(email, password string) (entity.User, error)
	Confirm(userId string) error
}

type userStorage struct {
	db *gorm.DB
}

func newUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

func (s *userStorage) Add(user entity.User) error {
	result := s.db.Create(&user)
	if result.Error != nil {
		return errors.New("користувач вже зареєстрований")
	}

	return nil
}

func (s *userStorage) GetByCredentials(email, password string) (entity.User, error) {
	var user entity.User

	result := s.db.Take(&user, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return entity.User{}, errors.New("користувача не знайдено")
	}

	return user, nil
}

func (s *userStorage) Confirm(userId string) error {
	var user entity.User

	result := s.db.Model(&user).Where("id = ?", userId).Update("is_confirmed", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
