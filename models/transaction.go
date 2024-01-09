package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model    `json:"-"`
	ID            uint `json:"id" gorm:"primarykey"`
	HouseID       uint `json:"house_id"`
	UserID        uint `json:"user_id"`
	ReservationID uint `json:"reservation_id"`
}
