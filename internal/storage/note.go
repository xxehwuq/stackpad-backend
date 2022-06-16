package storage

import (
	"errors"

	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type NoteStorage interface {
	Add(note entity.Note) error
	Update(note entity.Note) error
	GetAll(userId string) ([]entity.Note, error)
	GetAllFromNotebook(notebookId, userId string) ([]entity.Note, error)
	GetAllBookmarks(userId string) ([]entity.Note, error)
	GetById(noteId, userId string) (entity.Note, error)
	DeleteById(note entity.Note) error
}

type noteStorage struct {
	db *gorm.DB
}

func newNoteStorage(db *gorm.DB) *noteStorage {
	return &noteStorage{
		db: db,
	}
}

func (s *noteStorage) Add(note entity.Note) error {
	result := s.db.Create(&note)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *noteStorage) Update(note entity.Note) error {
	result := s.db.Save(&note)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *noteStorage) GetAll(userId string) ([]entity.Note, error) {
	var notes []entity.Note

	result := s.db.Where("user_id = ?", userId).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}

	return notes, nil
}

func (s *noteStorage) GetAllFromNotebook(notebookId, userId string) ([]entity.Note, error) {
	var notes []entity.Note

	result := s.db.Where("notebook_id = ? AND user_id = ?", notebookId, userId).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}

	return notes, nil
}

func (s *noteStorage) GetAllBookmarks(userId string) ([]entity.Note, error) {
	var bookmarks []entity.Note

	result := s.db.Where("is_bookmarked = true AND user_id = ?", userId).Find(&bookmarks)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookmarks, nil
}

func (s *noteStorage) GetById(noteId, userId string) (entity.Note, error) {
	var note entity.Note

	result := s.db.Take(&note, "id = ? AND user_id = ?", noteId, userId)
	if result.Error != nil {
		return entity.Note{}, errors.New("нотатку не знайдено")
	}

	return note, nil
}

func (s *noteStorage) DeleteById(note entity.Note) error {
	result := s.db.Delete(&note)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
