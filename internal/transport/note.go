package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
)

func (t *Transport) noteAdd(ctx *gin.Context) {
	var note entity.Note
	userId := t.getUserId(ctx)

	if err := ctx.BindJSON(&note); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.Note.Add(note, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (t *Transport) noteGetAll(ctx *gin.Context) {
	userId := t.getUserId(ctx)

	notes, err := t.service.Note.GetAll(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (t *Transport) noteGetById(ctx *gin.Context) {
	userId := t.getUserId(ctx)
	noteId := ctx.Param("id")

	note, err := t.service.Note.GetById(noteId, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, note)
}
