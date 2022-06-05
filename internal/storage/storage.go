package storage

import "gorm.io/gorm"

type Storage struct {
	User UserStorage
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		User: newUserStorage(db),
	}
}
