package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type GetCitiesResponse struct {
	StatusCode int            `json:"status_code"`
	Code       string         `json:"code"`
	Message    string         `json:"message"`
	Cities     []*models.City `json:"cities"`
}

func (_ *GetCitiesResponse) CreateCitiesResponse(status int, code, message string, cities []*models.City) *GetCitiesResponse {
	return &GetCitiesResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Cities:     cities,
	}
}
