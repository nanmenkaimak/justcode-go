package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/models"
	"net/http"
	"strconv"
)

func (m *Repository) UpdateUserByID(ctx *gin.Context) {
	var updateUser models.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser.ID = id

	if err := ctx.BindJSON(&updateUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, err := m.DB.UpdateUser(updateUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errors.New("no rows affected")})
		return
	}
	ctx.Status(http.StatusNoContent)
}
