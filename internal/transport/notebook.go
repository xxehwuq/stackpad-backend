package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
)

func (t *Transport) notebookAdd(ctx *gin.Context) {
	var notebook entity.Notebook
	userId := t.getUserId(ctx)

	if err := ctx.BindJSON(&notebook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.Notebook.Add(notebook, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (t *Transport) notebookGetAll(ctx *gin.Context) {
	userId := t.getUserId(ctx)

	notebooks, err := t.service.Notebook.GetAll(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, notebooks)
}

func (t *Transport) notebookGetById(ctx *gin.Context) {
	userId := t.getUserId(ctx)
	notebookId := ctx.Param("id")

	notebook, err := t.service.Notebook.GetById(notebookId, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, notebook)
}
