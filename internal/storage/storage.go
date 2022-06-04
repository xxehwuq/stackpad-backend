package storage

import "gorm.io/gorm"

type Storage struct {
	Notebook NotebookStorage
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		Notebook: newNotebookStorage(db),
	}
}
