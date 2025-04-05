package main

import (
	"time"

	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	server := gin.Default()
	server.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:4200"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	routes.RegisterUserRoutes(server)
	routes.RegisterEventRoutes(server)
	server.Run()
}
