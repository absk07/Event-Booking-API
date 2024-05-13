package main

import (
	"fmt"
	"net/http"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err,
		})
		return
	}
	err = event.Save()
	if err != nil {
		// fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message":    "New Event successfully created!",
	})
}

func main() {
	db.InitDB()

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
