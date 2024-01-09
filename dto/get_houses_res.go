package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"time"
)

type House struct {
	ID            uint      `json:"id"`
	HouseName     string    `json:"house_name"`
	UserID        uint      `json:"user_id"`
	PricePerNight int       `json:"price_per_night"`
	Description   string    `json:"description"`
	CityID        uint      `json:"city_id"`
	CityName      string    `json:"city_name"`
	MaxGuest      int       `json:"max_guest"`
	CreatedDate   time.Time `json:"created_date"`
	UpdateDate    time.Time `json:"updated_date"`
}

func (_ *House) CreateHouseResponse(h *models.House) *House {
	return &House{
		ID:            h.ID,
		HouseName:     h.HouseName,
		UserID:        h.UserID,
		PricePerNight: h.PricePerNight,
		Description:   h.Description,
		CityID:        h.CityID,
		CityName:      h.CityName,
		MaxGuest:      h.MaxGuest,
	}
}

type HousesResponse struct {
	StatusCode int        `json:"status_code"`
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	HousesData []*House   `json:"houses_data"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"end_date"`
}

func (_ *House) ShowHouses(status int, code, message string, houses []*House, s, e *time.Time) *HousesResponse {
	return &HousesResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		HousesData: houses,
		StartDate:  s,
		EndDate:    e,
	}
}
