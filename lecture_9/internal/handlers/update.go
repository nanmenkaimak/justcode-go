package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
	"net/http"
)

func (m *Repository) UpdateGistByID(ctx *gin.Context) {
	var updatedGist models.GistRequest

	if err := ctx.BindJSON(&updatedGist); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	updatedGist.Gist.ID = id
	err = m.DB.UpdateGist(updatedGist)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
