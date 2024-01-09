package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type FirstPhotoResponse struct {
	StatusCode int                `json:"status_code"`
	Code       string             `json:"code"`
	Message    string             `json:"message"`
	Photo      *models.HousePhoto `json:"photo"`
}

func (_ *FirstPhotoResponse) CreateFirstPhotoResponse(status int, code, message string, p *models.HousePhoto) *FirstPhotoResponse {
	return &FirstPhotoResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Photo:      p,
	}
}
