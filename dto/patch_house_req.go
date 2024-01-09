package dto

type UpdateHouseRequest struct {
	HouseName     string `json:"house_name"`
	PricePerNight int    `json:"price_per_night"`
	Description   string `json:"description"`
	MaxGuest      int    `json:"max_guest"`
}
