package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type BookingsResponse struct {
	StatusCode  int                 `json:"status_code"`
	Code        string              `json:"code"`
	Message     string              `json:"message"`
	Reservation *models.Reservation `json:"reservation_data"`
}

func (_ *BookingsResponse) CreateBookingResponse(status int, code, message string, r *models.Reservation) *BookingsResponse {
	return &BookingsResponse{
		StatusCode:  status,
		Code:        code,
		Message:     message,
		Reservation: r,
	}
}
