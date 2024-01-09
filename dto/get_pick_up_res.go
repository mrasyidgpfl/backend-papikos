package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type PickUpResponse struct {
	StatusCode  int            `json:"status_code"`
	Code        string         `json:"code"`
	Message     string         `json:"message"`
	Reservation *models.PickUp `json:"reservation_data"`
	Wallet      *models.Wallet `json:"wallet"`
}

func (_ *PickUpResponse) CreatePickUpResponse(status int, code, message string, p *models.PickUp, w *models.Wallet) *PickUpResponse {
	return &PickUpResponse{
		StatusCode:  status,
		Code:        code,
		Message:     message,
		Reservation: p,
		Wallet:      w,
	}
}
