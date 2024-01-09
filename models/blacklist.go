package models

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Token      string `json:"token"`
}
