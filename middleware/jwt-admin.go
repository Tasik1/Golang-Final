package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAdmin, adminExists := ctx.Get("isAdmin")
		if adminExists && isAdmin == true {
			ctx.Next()
			return
		}

		id := ctx.Param("user_id")
		if ctxID, ok := ctx.Get("userID"); ok {
			ctxIDStr := fmt.Sprintf("%v", ctxID)
			if ctxIDStr == id {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
	}
}
