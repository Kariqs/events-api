package initializers

import (
	"fmt"

	"github.com/Kariqs/events-api/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Event{}, &models.User{})
	fmt.Println("Database Synced Successfully!")
}
