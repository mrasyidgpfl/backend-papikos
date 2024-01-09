package models

import "gorm.io/gorm"

type GamesChances struct {
	gorm.Model `json:"-"`
	ID         uint `json:"id" gorm:"primarykey"`
	UserId     uint `json:"user_id"`
	Chance     int  `json:"chance"`
	History    int  `json:"history"`
}
