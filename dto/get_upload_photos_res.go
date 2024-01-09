package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type UploadPhotosResponse struct {
	StatusCode int                  `json:"status_code"`
	Code       string               `json:"code"`
	Message    string               `json:"message"`
	Photos     []*models.HousePhoto `json:"photos"`
}

func (_ *UploadPhotosResponse) CreateUploadPhotosResponse(status int, code, message string, ps []*models.HousePhoto) *UploadPhotosResponse {
	return &UploadPhotosResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Photos:     ps,
	}
}
