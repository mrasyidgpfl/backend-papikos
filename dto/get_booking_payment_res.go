package dto

import "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"

type BookingPaymentResponse struct {
	StatusCode          int             `json:"status_code"`
	Code                string          `json:"code"`
	Message             string          `json:"message"`
	BookingsPaymentData *PaymentReceipt `json:"payment_data"`
}

func (_ *BookingPaymentResponse) CreateBookingPaymentResponse(status int, code, message string, receipt *PaymentReceipt) *BookingPaymentResponse {
	return &BookingPaymentResponse{
		StatusCode:          status,
		Code:                code,
		Message:             message,
		BookingsPaymentData: receipt,
	}
}

type PaymentReceipt struct {
	Transaction *models.Transaction `json:"transaction"`
	Wallet      *models.Wallet      `json:"wallet"`
	BookingCode uint                `json:"booking_code"`
	BookingInfo *models.Reservation `json:"booking_info"`
}
