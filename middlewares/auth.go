package middlewares

import (
	// "fmt"
	"net/http"

	"example.com/event-booking-api/db"
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
	id := ctx.Param("id")
	query := `SELECT userId FROM events WHERE id = $1`
	row := db.DB.QueryRow(query, id)
	var userId string
	err := row.Scan(&userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User Unauthorized!",
		})
		return
	}
	// fmt.Println(userId, ctx.GetString("userId"))
	if userId != ctx.GetString("userId") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User Unauthorized!",
		})
		return
	}
	ctx.Next()
}