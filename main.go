package main

import (
	"net/http"

	"github.com/Kariqs/events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(ctx *gin.Context) {
	client, context, cancel := models.DatabaseConnection()
	defer cancel()
	defer client.Disconnect(context)

	events := models.GetAllEvents(client)
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	client, context, cancel := models.DatabaseConnection()
	defer cancel()
	defer client.Disconnect(context)

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	event.UserID = 1
	savedEvent, err := event.Save(client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully.", "event": savedEvent})
}
