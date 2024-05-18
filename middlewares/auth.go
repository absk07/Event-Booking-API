package middlewares

import (
	"net/http"

	"example.com/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User Unauthorized!",
		})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User Unauthorized!",
		})
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}

func IsAuthor(ctx *gin.Context) {
	
}