package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ImageUrl    string         `json:"imageUrl" binding:"required"`
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Location    string         `json:"location" binding:"required"`
	Price       string         `json:"price" binding:"required"`
	Date        time.Time      `json:"date" binding:"required"`
	Tags        datatypes.JSON `json:"tags" binding:"required"`
}
