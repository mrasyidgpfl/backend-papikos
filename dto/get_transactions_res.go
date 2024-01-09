package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type GetTransactions struct {
	StatusCode   int                   `json:"status_code"`
	Code         string                `json:"code"`
	Message      string                `json:"message"`
	Transactions []*models.Transaction `json:"transactions"`
}

func (_ *GetTransactions) CreateTransactionsResponse(status int, code, message string, ts []*models.Transaction) *GetTransactions {
	return &GetTransactions{
		StatusCode:   status,
		Code:         code,
		Message:      message,
		Transactions: ts,
	}
}
