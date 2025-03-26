package main

import (
	"github.com/Kariqs/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
}
