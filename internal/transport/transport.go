package transport

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type Transport struct {
	service *service.Service
	pkg     entity.Pkg
}

func New(service *service.Service, pkg entity.Pkg) *Transport {
	return &Transport{
		service: service,
		pkg:     pkg,
	}
}

func (t *Transport) Init(cfg *config.Config) {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://stackpad.vercel.app"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin"},
	}))
	router.Use(favicon.New("favicon.ico"))

	router.SetTrustedProxies([]string{"https://stackpad.vercel.app"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.ClientIP())
	})
	t.initApi(router)

	log.Fatal(router.Run())
	// log.Fatal(router.Run(":" + cfg.Http.Port))
}

func (t *Transport) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		authApi := api.Group("", t.setUserId)

		notebook := authApi.Group("/notebook")
		{
			notebook.POST("", t.notebookAdd)
			notebook.GET("", t.notebookGetAll)
			notebook.GET("/:id", t.notebookGetById)
		}

		note := authApi.Group("/note")
		{
			note.POST("", t.noteAdd)
			note.PUT("", t.noteUpdate)
			note.GET("/notebook/:notebookId", t.noteGetAllFromNotebook)
			note.GET("/bookmarks", t.noteGetAllBookmarks)
			note.GET("/:id", t.noteGetById)
			note.DELETE("/:id", t.noteDeleteById)
		}

		user := api.Group("/user")
		{
			user.POST("/sign-up", t.userSignUp)
			user.POST("/sign-in", t.userSignIn)
			authApi.GET("/user/confirm", t.userConfirm)
		}
	}
}
