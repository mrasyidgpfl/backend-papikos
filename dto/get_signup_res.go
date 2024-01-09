package dto

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"gorm.io/gorm"
)

type SignUpResponse struct {
	StatusCode   int                 `json:"status_code"`
	Code         string              `json:"code"`
	Message      string              `json:"message"`
	User         models.User         `json:"data"`
	Wallet       models.Wallet       `json:"wallet"`
	GamesChances models.GamesChances `json:"games_chances"`
}

func (_ *SignUpResponse) CreateSignUpResponse(status int, code, message string, u *models.User, w *models.Wallet, gC *models.GamesChances) *SignUpResponse {
	return &SignUpResponse{
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
		Wallet: models.Wallet{
			Model:   gorm.Model{},
			ID:      w.ID,
			UserId:  u.ID,
			Balance: 0,
		},
		GamesChances: models.GamesChances{
			Model:  gorm.Model{},
			ID:     gC.ID,
			UserId: u.ID,
			Chance: 0,
		},
	}
}
