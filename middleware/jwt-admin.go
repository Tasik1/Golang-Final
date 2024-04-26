package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAdmin, exists := ctx.Get("isAdmin")
		if !exists || isAdmin != true {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}
		ctx.Next()
	}
}
