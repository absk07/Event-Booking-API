package routes

import (
	"example.com/event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// events routes
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", middlewares.IsAuthenticated, createEvent)
	router.PUT("/event/:id/edit", middlewares.IsAuthenticated, editEvent)
	router.DELETE("/event/:id/delete", middlewares.IsAuthenticated, deleteEvent)

	/** OR
		authenticated := router.Group("/")
		authenticated.Use(middlewares.IsAuthenticated)
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/event/:id/edit", editEvent)
		authenticated.DELETE("/event/:id/delete", deleteEvent)
	*/

	// users routes
	router.POST("/signup", signup)
	router.POST("/signin", signin)
}