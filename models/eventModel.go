package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ImageUrl   string
	Title      string
	Desription string
	Price      string
	Date       time.Time
}
