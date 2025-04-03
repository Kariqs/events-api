package routes

import (
	"github.com/Kariqs/events-api/controllers"
	"github.com/Kariqs/events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", middlewares.RequireAuth(), controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events/:id", controllers.GetEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
}
