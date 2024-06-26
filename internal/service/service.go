package service

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
)

type Service struct {
	User     UserService
	Notebook NotebookService
	Note     NoteService
}

func New(storage *storage.Storage, pkg entity.Pkg) *Service {
	return &Service{
		User:     newUserService(storage.User, pkg.PasswordManager, pkg.JwtManager, pkg.MailManager),
		Notebook: newNotebookService(storage.Notebook),
		Note:     newNoteService(storage.Note),
	}
}
