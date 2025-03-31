package initializers

import "github.com/Kariqs/events-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Event{}, &models.User{})
}
