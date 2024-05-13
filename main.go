package main

import (
	"net/http"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Something went wrong!",
		})
		return
	}
	event.Id = uuid.New().String()
	event.UserId = uuid.New().String()
	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}

func main() {
	server := gin.Default()

	server.GET("/health-checker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Server OK!",
		})
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":3000")
}
