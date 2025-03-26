package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name" binding:"required"`
	Description string             `bson:"description" binding:"required"`
	Location    string             `bson:"location" binding:"required"`
	Date        time.Time          `bson:"date" binding:"required"`
	UserID      int
}

var events = []Event{}

func (event *Event) Save(client *mongo.Client) (*mongo.InsertOneResult, error) {
	collection := client.Database("eventsdb").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, event)
	if err != nil {
		log.Println("Failed to insert event:", err)
		return nil, err
	}
	return result, nil
}

func GetAllEvents() []Event {
	return events
}
