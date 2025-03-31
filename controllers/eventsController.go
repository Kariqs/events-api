package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEvent(ctx *gin.Context) {
	//Get data off request body
	var event models.Event
	ctx.ShouldBindJSON(&event)
	//Create an event
	result := initializers.DB.Create(&event)

	if result.Error != nil {
		log.Fatal("Unable to create event.")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to create event.", "status": http.StatusBadRequest})
	}
	//Return the created event
	ctx.JSON(http.StatusCreated, gin.H{"event": event})
}

func GetEvents(ctx *gin.Context) {
	var events []models.Event
	result := initializers.DB.Find(&events)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to fetch events.", "status": http.StatusBadRequest})
	}
	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
	//Get Id from the URL
	eventId := ctx.Param("id")

	//Find event with the ID
	var event models.Event
	result := initializers.DB.First(&event, eventId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Database error"})
		return
	}

	//Return the event
	ctx.JSON(http.StatusOK, event)
}

func UpdateEvent(ctx *gin.Context) {
	//Get Id and data from the request
	var updateData models.Event
	eventId := ctx.Param("id")
	ctx.ShouldBindJSON(&updateData)
	//Find event with the ID
	var event models.Event
	result := initializers.DB.First(&event, eventId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	//Update the event
	initializers.DB.Model(&event).Updates(updateData)

	//Return the event
	ctx.JSON(http.StatusOK, event)
}

func DeleteEvent(ctx *gin.Context) {
	//Get Id from URL
	eventId := ctx.Param("id")
	var event models.Event

	//Check if event exists
	result := initializers.DB.First(&event, eventId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	//Delete event
	deleteResult := initializers.DB.Delete(&event)
	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unabale to delete event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event has been deleted", "deletedEvent": deleteResult})
}
