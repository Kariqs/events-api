package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEventById)
	server.DELETE("/events/:eventId", deleteEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}
