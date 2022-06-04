package transport

import (
	"fmt"

	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type NotebookTransport interface {
	Test()
}

type notebookTransport struct {
	service service.NotebookService
}

func newNotebookTransport(service service.NotebookService) *notebookTransport {
	return &notebookTransport{
		service: service,
	}
}

func (s *notebookTransport) Test() {
	fmt.Println("Test() from notebook transport")
	s.service.Test()
}
