package models

import "gorm.io/gorm"

type City struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
}
