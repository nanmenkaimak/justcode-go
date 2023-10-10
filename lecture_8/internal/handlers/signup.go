package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/models"
	"net/http"
)

func (m *Repository) SignUp(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := m.DB.CreateUser(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}
