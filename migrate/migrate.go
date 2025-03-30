package main

import (
	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Event{})
}
