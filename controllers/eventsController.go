package controllers

import (
	"log"
	"net/http"

	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/models"
	"github.com/gin-gonic/gin"
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
