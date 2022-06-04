package app

import (
	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
	"github.com/yaroslavyarosh/stackpad-backend/internal/storage"
	"github.com/yaroslavyarosh/stackpad-backend/internal/transport"
)

func Run(cfg *config.Config) {
	storage := storage.New()
	service := service.New(storage)
	transport := transport.New(service)

	transport.Notebook.Test()
}
