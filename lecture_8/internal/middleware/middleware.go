package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const authToken = "tokenmokenkoken"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth header"})
			return
		}
		if header != authToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong token"})
			return
		}
		ctx.Next()
	}
}
