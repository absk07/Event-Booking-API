package routes

import (
	"fmt"
	"net/http"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
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
			"error":   err,
		})
		return
	}
	// fmt.Println("user id", userId)
	event.UserId = ctx.GetString("userId")
	err = event.Save()
	if err != nil {
		// fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "New Event successfully created!",
	})
}

func getEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := models.GetEventById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func editEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := models.GetEventById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	updatedEvent.Id = id
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "Event updated successfully!",
	})
}

func deleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "Event deleted successfully!",
	})
}
