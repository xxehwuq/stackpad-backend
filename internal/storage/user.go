package storage

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type UserStorage interface {
	Add(user entity.User) error
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
		return result.Error
	}

	return nil
}
