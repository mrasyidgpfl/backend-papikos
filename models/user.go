package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"full_name"`
	Address    string `json:"address"`
	CityId     uint   `json:"city_id"`
	Role       string `json:"role"`
}
