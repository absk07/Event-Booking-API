package routes

import (
	// "fmt"
	"net/http"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	eventId := ctx.Param("id")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Event not found!",
		})
		return
	}
	err = event.RegisterInEvent(userId)
	// fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Could not register user for event!",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"error":   "Registered in event!",
	})
}

func cancleRegistration(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	eventId := ctx.Param("id")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Event not found!",
		})
		return
	}
	err = event.CancleRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Could not cancle user registartion for event!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "Event registration cancled!",
	})
}