package routes

import (
	"github.com/Kariqs/events-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", controllers.CreateEvent)
	server.Run()
}
