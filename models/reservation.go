package models

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model   `json:"-"`
	ID           uint      `json:"id" gorm:"primarykey"`
	HouseID      uint      `json:"house_id"`
	UserID       uint      `json:"user_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	TotalPrice   int       `json:"total_price"`
}
