package services

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"net/http"
)

type WalletService interface {
	TopUp(d *dto.TopUpRequest) (*dto.TopUpResponse, error)
	GetWalletFromUserId(id uint) (*models.Wallet, error)
	Pay(userId uint, input *dto.BookingPaymentRequest) (*dto.BookingPaymentResponse, error)
}

type walletService struct {
	walletRepository repositories.WalletRepository
	appConfig        config.AppConfig
}

type WalletConfig struct {
	WalletRepository repositories.WalletRepository
	AppConfig        config.AppConfig
}

func NewWalletService(c *WalletConfig) WalletService {
	return &walletService{
		walletRepository: c.WalletRepository,
		appConfig:        c.AppConfig,
	}
}

func (w *walletService) GetWalletFromUserId(id uint) (*models.Wallet, error) {
	walletId, err := w.walletRepository.FindWalletFromUserId(id)
	if walletId == nil || err != nil {
		return nil, err
	}
	return walletId, nil
}

func (w *walletService) TopUp(req *dto.TopUpRequest) (*dto.TopUpResponse, error) {
	topUp := dto.TopUpDetail{
		UserID:   req.UserID,
		Wallet:   req.Wallet,
		Amount:   req.Amount,
		SourceID: req.SourceID,
	}
	topUpPtr, err := w.walletRepository.TopUp(req)
	if topUpPtr == false {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED ERROR",
			Message:    "Unauthorized error.",
		}
	}
	return new(dto.TopUpResponse).CreateTopUpRes(200, "TOP UP SUCCESS", "Top up success", topUp), err
}

func (w *walletService) Pay(userId uint, input *dto.BookingPaymentRequest) (*dto.BookingPaymentResponse, error) {
	transactionPaid, _ := w.walletRepository.CheckTransaction(input.ReservationID)
	if transactionPaid {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "ALREADY PAID ERROR",
			Message:    "House reservation has already been paid",
		}
	}
	bookingInfo, _ := w.walletRepository.FindBooking(input.ReservationID)
	transactionId, _ := w.GenerateTransactionId()
	checkBalance, _ := w.walletRepository.CheckBalance(userId, bookingInfo.TotalPrice)
	if !checkBalance {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "INSUFFICIENT BALANCE ERROR",
			Message:    "Balance insufficient, please top up",
		}
	}
	transaction, wallet, paymentErr := w.walletRepository.PayBooking(userId, transactionId, input.ReservationID, bookingInfo.TotalPrice)
	if paymentErr != nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusNotFound,
			Code:       "NOT FOUND ERROR",
			Message:    "Not found error and payment was unsuccessful",
		}
	}

	response := dto.PaymentReceipt{
		Transaction: transaction,
		Wallet:      wallet,
		BookingCode: input.ReservationID,
		BookingInfo: bookingInfo,
	}

	return new(dto.BookingPaymentResponse).CreateBookingPaymentResponse(200, "BOOKING PAYMENT SUCCESSFUL", "Payment for reservation was successful", &response), paymentErr
}

func (w *walletService) GenerateTransactionId() (uint, error) {
	totalTransactions, _ := w.walletRepository.CountTotalTransactions()
	totalTransactions += 1
	transactionId := uint(600000 + totalTransactions)
	return transactionId, nil
}
