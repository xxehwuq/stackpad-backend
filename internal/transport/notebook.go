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
