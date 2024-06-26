package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroslavyarosh/stackpad-backend/internal/entity"
)

func (t *Transport) noteAdd(ctx *gin.Context) {
	var note entity.Note

	note.UserId = t.getUserId(ctx)

	if err := ctx.BindJSON(&note); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.Note.Add(note)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (t *Transport) noteUpdate(ctx *gin.Context) {
	var note entity.Note

	note.UserId = t.getUserId(ctx)

	if err := ctx.BindJSON(&note); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := t.service.Note.Update(note)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "")
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

func (t *Transport) noteGetAllFromNotebook(ctx *gin.Context) {
	userId := t.getUserId(ctx)
	notebookId := ctx.Param("notebookId")

	notes, err := t.service.Note.GetAllFromNotebook(notebookId, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (t *Transport) noteGetAllBookmarks(ctx *gin.Context) {
	userId := t.getUserId(ctx)

	bookmarks, err := t.service.Note.GetAllBookmarks(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bookmarks)
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

func (t *Transport) noteDeleteById(ctx *gin.Context) {
	var note entity.Note

	note.UserId = t.getUserId(ctx)
	note.Id = ctx.Param("id")

	err := t.service.Note.DeleteById(note)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "")
}
