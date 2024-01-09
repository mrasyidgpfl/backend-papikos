package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type GetPickUps struct {
	StatusCode int      `json:"status_code"`
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	PickUps    []*PURes `json:"pickups"`
}

type PURes struct {
	PickUp *models.PickUp `json:"pickup"`
	Status string         `json:"status"`
}

func (_ *GetPickUps) GetPickUps(statusCode int, code, message string, PURes []*PURes) *GetPickUps {
	return &GetPickUps{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		PickUps:    PURes,
	}
}
