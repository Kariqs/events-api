package routes

import (
	"github.com/Kariqs/events-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events/:id", controllers.GetEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
	server.Run()
}
