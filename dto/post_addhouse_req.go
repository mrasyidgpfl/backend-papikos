package dto

type AddHouseRequest struct {
	HouseName     string `json:"house_name"`
	PricePerNight int    `json:"price_per_night"`
	Description   string `json:"description"`
	CityID        uint   `json:"city_id"`
	CityName      string `json:"city_name"`
	MaxGuest      int    `json:"max_guest"`
}
