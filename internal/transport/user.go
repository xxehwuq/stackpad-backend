package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
)

func (t *Transport) userSignUp(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := t.service.User.SignUp(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (t *Transport) userSignIn(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := t.service.User.SignIn(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (t *Transport) userConfirm(ctx *gin.Context) {
	userId := ctx.Param("userId")

	err := t.service.User.Confirm(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
