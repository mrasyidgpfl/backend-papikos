package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type ShowBookingsResponse struct {
	StatusCode   int                   `json:"status_code"`
	Code         string                `json:"code"`
	Message      string                `json:"message"`
	BookingsData []*models.Reservation `json:"bookings_data"`
}

func (_ *ShowBookingsResponse) ShowBookings(status int, code, message string, bookings []*models.Reservation) *ShowBookingsResponse {
	return &ShowBookingsResponse{
		StatusCode:   status,
		Code:         code,
		Message:      message,
		BookingsData: bookings,
	}
}
