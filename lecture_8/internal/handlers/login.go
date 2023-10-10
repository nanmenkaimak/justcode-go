package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *Repository) Login(ctx *gin.Context) {
	var user LoginParams

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := m.DB.Authenticate(user.Email, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
