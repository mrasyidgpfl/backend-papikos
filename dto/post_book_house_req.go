package dto

import "time"

type BookHouseRequest struct {
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
}
