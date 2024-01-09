package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type GetHouseResponse struct {
	StatusCode int           `json:"status_code"`
	Code       string        `json:"code"`
	Message    string        `json:"message"`
	House      *models.House `json:"house"`
}

func (_ *GetHouseResponse) CreateGetHouseResponse(status int, code, message string, h *models.House) *GetHouseResponse {
	return &GetHouseResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		House:      h,
	}
}
