package main

import (
	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
}
