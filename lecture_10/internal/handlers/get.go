package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
	"net/http"
	"time"
)

func (m *Repository) GetAllGists(ctx *gin.Context) {
	res, err := m.DBRedis.Get(context.Background(), "all").Result()
	if err == nil {
		var gist []models.GistRequest
		json.Unmarshal([]byte(res), &gist)
		ctx.JSON(http.StatusOK, gist)
		return
	}
	allGists, err := m.DB.GetAllGists()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonGist, err := json.Marshal(allGists)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = m.DBRedis.Set(context.Background(), "all", jsonGist, 1*time.Minute).Err()
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
	res, err := m.DBRedis.Get(context.Background(), id.String()).Result()
	if err == nil {
		var gist models.GistRequest
		json.Unmarshal([]byte(res), &gist)
		ctx.JSON(http.StatusOK, gist)
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

	jsonGist, err := json.Marshal(gist)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = m.DBRedis.Set(context.Background(), id.String(), jsonGist, 1*time.Minute).Err()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gist)
}
