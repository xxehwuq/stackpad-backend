package transport

import (
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type Transport struct {
	Notebook NotebookTransport
}

func New(service *service.Service) *Transport {
	return &Transport{
		Notebook: newNotebookTransport(service.Notebook),
	}
}

func (t *Transport) Init() {

}
