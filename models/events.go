package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" binding:"required" json:"name"`
	Description string             `bson:"description" binding:"required" json:"description"`
	Location    string             `bson:"location" binding:"required" json:"location"`
	Date        time.Time          `bson:"date" binding:"required" json:"date"`
	UserID      int
}

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

func GetAllEvents(client *mongo.Client) ([]map[string]any, error) {
	collection := client.Database("eventsdb").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer cur.Close(ctx)

	var events []map[string]any

	for cur.Next(ctx) {
		var event map[string]any
		err := cur.Decode(&event)
		if err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		events = append(events, event)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return events, nil
}

func GetEventById(client *mongo.Client, eventId string) (map[string]any, error) {
	collection := client.Database("eventsdb").Collection("events")

	objectId, err := bson.ObjectIDFromHex(eventId)
	if err != nil {
		return nil, fmt.Errorf("invalid event ID format: %v", err)
	}

	var event map[string]any
	result := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})
	err = result.Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no event found with ID %s", eventId)
		}
		return nil, fmt.Errorf("error finding event: %v", err)
	}
	return event, nil
}
