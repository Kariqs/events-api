package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DatabaseConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading environment variable failed.")
	}

	mongoDbUrl := os.Getenv("MONGODB_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(options.Client().ApplyURI(mongoDbUrl))

	if err != nil {
		fmt.Println("error connecting to MongoDB:", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err.Error())
	}

	fmt.Println("Successfully connected to MongoDB")

	return client, ctx, cancel
}
