package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type TopUpResponse struct {
	StatusCode  int         `json:"status_code"`
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	TopUpDetail TopUpDetail `json:"data"`
}

func (_ *TopUpResponse) CreateTopUpRes(status int, code, message string, t TopUpDetail) *TopUpResponse {
	return &TopUpResponse{
		StatusCode: status,
		Code:       code,
		Message:    message,
		TopUpDetail: TopUpDetail{
			UserID:   t.UserID,
			Wallet:   t.Wallet,
			Amount:   t.Amount,
			SourceID: t.SourceID,
		},
	}
}

type TopUpDetail struct {
	UserID   uint           `json:"user_id"`
	Wallet   *models.Wallet `json:"wallet"`
	Amount   int            `json:"amount"`
	SourceID uint           `json:"source_id"`
}
