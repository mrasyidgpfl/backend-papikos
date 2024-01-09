package models

import (
	"gorm.io/gorm"
	"time"
)

type House struct {
	gorm.Model    `json:"-"`
	ID            uint      `json:"id" gorm:"primarykey"`
	HouseName     string    `json:"house_name"`
	UserID        uint      `json:"user_id"`
	PricePerNight int       `json:"price_per_night"`
	Description   string    `json:"description"`
	CityID        uint      `json:"city_id"`
	CityName      string    `json:"city_name"`
	MaxGuest      int       `json:"max_guest"`
	CreatedAt     time.Time `json:"created_at"`
}
