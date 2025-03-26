package routes

import (
	"fmt"
	"net/http"

	"github.com/Kariqs/events-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	client, context, cancel := models.DatabaseConnection()
	defer cancel()
	defer client.Disconnect(context)

	events, err := models.GetAllEvents(client)

	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
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

func getEventById(ctx *gin.Context) {
	client, context, cancel := models.DatabaseConnection()
	defer cancel()
	defer client.Disconnect(context)

	eventId := ctx.Param("eventId")
	fmt.Println(eventId)
	event, err := models.GetEventById(client, eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
