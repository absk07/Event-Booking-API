package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// events routes
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", createEvent)
	router.PUT("/event/:id/edit", editEvent)
	router.DELETE("/event/:id/delete", deleteEvent)

	// users routes
	
}