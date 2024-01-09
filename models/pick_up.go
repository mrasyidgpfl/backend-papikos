package models

import (
	"gorm.io/gorm"
)

type PickUp struct {
	gorm.Model     `json:"-"`
	ID             uint `json:"id" gorm:"primarykey"`
	UserID         uint `json:"user_id"`
	ReservationID  uint `json:"reservation_id"`
	PickUpStatusID uint `json:"pick_up_status_id"`
}
