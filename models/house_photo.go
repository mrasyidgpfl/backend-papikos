package models

import "gorm.io/gorm"

type HousePhoto struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	HouseID    uint   `json:"house_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhotoURL   string `json:"photo_url"`
}
