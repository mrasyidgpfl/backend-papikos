package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"time"
)

type UpdateHouseResponse struct {
	StatusCode int           `json:"status_code"`
	Code       string        `json:"code"`
	Message    string        `json:"message"`
	House      *models.House `json:"house"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

func (_ *UpdateHouseResponse) CreateUpdateHouseResponse(status int, code, message string, h *models.House, t time.Time) *UpdateHouseResponse {
	return &UpdateHouseResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		House:      h,
		UpdatedAt:  t,
	}
}
