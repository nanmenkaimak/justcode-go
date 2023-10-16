package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (m *Repository) GetAllGists(ctx *gin.Context) {
	allGists, err := m.DB.GetAllGists()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, allGists)
}

func (m *Repository) GetGistByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid uuid").Error()})
		return
	}

	gist, err := m.DB.GetGistByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if gist.Gist.ID == uuid.Nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errors.New("no gist with such id").Error()})
		return
	}

	ctx.JSON(http.StatusOK, gist)
}
