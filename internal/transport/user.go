package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type UserTransport interface {
	SignUp(ctx *gin.Context)
}

type userTransport struct {
	service service.UserService
}

func newUserTransport(service service.UserService) *userTransport {
	return &userTransport{
		service: service,
	}
}

func (s *userTransport) SignUp(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.service.SignUp(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
