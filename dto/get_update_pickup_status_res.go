package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type UpdatePickUpStatusResponse struct {
	StatusCode int                  `json:"status_code"`
	Code       string               `json:"code"`
	Message    string               `json:"message"`
	Status     *models.PickUpStatus `json:"status"`
}

func (_ *UpdatePickUpStatusResponse) CreateUpdatePickUpStatusRes(status int, code, message string, p *models.PickUpStatus) *UpdatePickUpStatusResponse {
	return &UpdatePickUpStatusResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Status:     p,
	}
}
