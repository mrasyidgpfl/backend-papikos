package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type UpdateProfileResponse struct {
	StatusCode int         `json:"status_code"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	User       models.User `json:"data"`
}

func (_ *UpdateProfileResponse) CreateUpdateProfileRes(status int, code, message string, u *models.User) *UpdateProfileResponse {
	return &UpdateProfileResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		User: models.User{
			ID:       u.ID,
			Email:    u.Email,
			FullName: u.FullName,
			Address:  u.Address,
			CityId:   u.CityId,
			Role:     u.Role,
		},
	}
}
