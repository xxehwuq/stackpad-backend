package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/config"
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

func (t *Transport) Init(cfg *config.Config) {
	router := gin.Default()
	t.initApi(router)

	router.Run(fmt.Sprintf(":%s", cfg.Http.Port))
}

func (t *Transport) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		notebooks := api.Group("/notebooks")
		{
			notebooks.GET("", t.Notebook.Test)
		}
	}
}
