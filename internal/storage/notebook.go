package storage

import "fmt"

type NotebookStorage interface {
	Test()
}

type notebookStorage struct {
}

func newNotebookStorage() *notebookStorage {
	return &notebookStorage{}
}

func (s *notebookStorage) Test() {
	fmt.Println("Test() from notebook storage")
}
