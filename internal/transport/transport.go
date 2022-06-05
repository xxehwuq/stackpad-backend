package transport

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type Transport struct {
	User UserTransport
}

func New(service *service.Service) *Transport {
	return &Transport{
		User: newUserTransport(service.User),
	}
}

func (t *Transport) Init(cfg *config.Config) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}))

	t.initApi(router)

	router.Run(fmt.Sprintf(":%s", cfg.Http.Port))
}

func (t *Transport) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/sign-up", t.User.SignUp)
			user.POST("/sign-in", t.User.SignIn)
		}
	}
}
