package main

import (
	"net/http"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/health-checker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Server OK!",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":3000")
}
