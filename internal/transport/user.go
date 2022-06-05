package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
	"github.com/yaroslavyarosh/stackpad-backend/internal/service"
)

type UserTransport interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	Confirm(ctx *gin.Context)
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

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (s *userTransport) SignIn(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.service.SignIn(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (s *userTransport) Confirm(ctx *gin.Context) {
	userId := ctx.Param("userId")

	err := s.service.Confirm(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
