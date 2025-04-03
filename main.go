package main

import (
	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	server := gin.Default()
	routes.RegisterUserRoutes(server)
	routes.RegisterEventRoutes(server)
	server.Run()
}
