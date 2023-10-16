package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
	"net/http"
)

func (m *Repository) CreateGist(ctx *gin.Context) {
	var newGist models.GistRequest

	if err := ctx.BindJSON(&newGist); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := m.DB.CreateGist(newGist)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
