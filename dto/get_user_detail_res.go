package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

type UserDetail struct {
	ID           uint                 `json:"id"`
	Email        string               `json:"email"`
	FullName     string               `json:"full_name"`
	Address      string               `json:"address"`
	CityID       uint                 `json:"city_id"`
	Role         string               `json:"role"`
	Wallet       *models.Wallet       `json:"wallet"`
	GamesChances *models.GamesChances `json:"games_chances"`
}

type UserDetailResponse struct {
	StatusCode int        `json:"status_code"`
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	UserDetail UserDetail `json:"user_data"`
}

func (_ *UserDetailResponse) CreateUserDetailRes(status int, code, message string, u *models.User, w *models.Wallet, gC *models.GamesChances) *UserDetailResponse {
	return &UserDetailResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		UserDetail: UserDetail{
			ID:           u.ID,
			Email:        u.Email,
			FullName:     u.FullName,
			Address:      u.Address,
			CityID:       u.CityId,
			Role:         u.Role,
			Wallet:       w,
			GamesChances: gC,
		},
	}
}
