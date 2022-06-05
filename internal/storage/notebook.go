package storage

import (
	"fmt"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"gorm.io/gorm"
)

type NotebookStorage interface {
	Test()
}

type notebookStorage struct {
	db *gorm.DB
}

func newNotebookStorage(db *gorm.DB) *notebookStorage {
	return &notebookStorage{
		db: db,
	}
}

func (s *notebookStorage) Test() {
	id, _ := gonanoid.New()
	// notebook := entity.Notebook{Id: id, Title: "Some title"}

	if result := s.db.Create(&entity.User{Id: id, Name: "Yaroslav", Email: "yaroslavyarosh.com@protonmail.com", Password: "LT4hlIGDzp1SspZ8ev6B"}); result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Test() from notebook storage")
}
