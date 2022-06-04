package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type NotebookTransport interface {
	Test(ctx *gin.Context)
}

type notebookTransport struct {
	service service.NotebookService
}

func newNotebookTransport(service service.NotebookService) *notebookTransport {
	return &notebookTransport{
		service: service,
	}
}

func (s *notebookTransport) Test(ctx *gin.Context) {
	fmt.Println("Test() from notebook transport")
	s.service.Test()
}
