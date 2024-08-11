package domain

import (
	"time"

	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model
	Id     int       `json:"id"`
	Time   time.Time `json:"time"`
	Place  Place     `json:"place" gorm:"embedded"`
	QrCode string    `json:"qrcode"`
}

type Place struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
