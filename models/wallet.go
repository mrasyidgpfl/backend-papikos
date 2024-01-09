package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model `json:"-"`
	ID         uint `json:"id" gorm:"primarykey"`
	UserId     uint `json:"user_id"`
	Balance    int  `json:"balance"`
}
