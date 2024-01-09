package models

import (
	"gorm.io/gorm"
)

type PickUpStatus struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Status     string `json:"status"`
}
