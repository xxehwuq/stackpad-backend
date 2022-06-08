package service

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type NoteService interface {
	Add(note entity.Note, userId string) (string, error)
	GetAll(userId string) ([]entity.Note, error)
	GetById(noteId, userId string) (entity.Note, error)
}

type noteService struct {
	storage storage.NoteStorage
}

func newNoteService(storage storage.NoteStorage) *noteService {
	return &noteService{
		storage: storage,
	}
}

func (s *noteService) Add(note entity.Note, userId string) (string, error) {
	id, _ := gonanoid.New()

	note.Id = id
	note.UserId = userId

	err := s.storage.Add(note)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *noteService) GetAll(userId string) ([]entity.Note, error) {
	notes, err := s.storage.GetAll(userId)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *noteService) GetById(noteId, userId string) (entity.Note, error) {
	note, err := s.storage.GetById(noteId, userId)
	if err != nil {
		return entity.Note{}, err
	}

	return note, nil
}
