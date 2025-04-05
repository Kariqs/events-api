package main

import (
	"fmt"
	"log"

	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Event{}, &models.User{})
	if err != nil {
		log.Fatal("Migration err:", err)
		return
	}
	fmt.Println("Migration was successful!")
}
