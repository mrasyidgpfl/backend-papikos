package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type FlipCoinsResponse struct {
	StatusCode    int            `json:"status_code"`
	Code          string         `json:"code"`
	Message       string         `json:"message"`
	FlipCoinsData FlipCoinsData  `json:"flip_coins_data"`
	Wallet        *models.Wallet `json:"wallet"`
}

func (_ *FlipCoinsResponse) CreateCoinsResponse(status int, code, message string, fCD FlipCoinsData, w *models.Wallet) *FlipCoinsResponse {
	return &FlipCoinsResponse{
		StatusCode:    status,
		Code:          code,
		Message:       message,
		FlipCoinsData: fCD,
		Wallet:        w,
	}
}

type FlipCoinsData struct {
	Chance     int    `json:"chances_left"`
	CoinSide   string `json:"coin_side"`
	GameResult string `json:"game_result"`
	Prize      int    `json:"prize"`
}
