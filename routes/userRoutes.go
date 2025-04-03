package routes

import (
	"github.com/Kariqs/events-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/auth/signup", controllers.SignUp)
	server.POST("/auth/login", controllers.Login)
}
