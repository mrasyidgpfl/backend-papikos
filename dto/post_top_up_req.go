package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type TopUpRequest struct {
	UserID   uint           `json:"user_id"`
	Wallet   *models.Wallet `json:"wallet"`
	Amount   int            `json:"amount"`
	SourceID uint           `json:"source_id"`
}
