package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.SetTrustedProxies([]string{"http://192.168.88.252:3000"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.ClientIP())
	})

	// router.Use(cors.New(cors.DefaultConfig()))
	router.Use(CORSMiddleware())

	t.initApi(router)

	log.Fatal(router.Run(":" + cfg.Http.Port))
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
