package repositories

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
)

type WalletRepository interface {
	FindWalletFromUserId(id uint) (*models.Wallet, error)
	TopUp(d *dto.TopUpRequest) (bool, error)
	CheckTransaction(bookingId uint) (bool, error)
	CountTotalTransactions() (int, error)
	CheckBalance(userId uint, price int) (bool, error)
	PayBooking(userId, transactionId, bookingId uint, price int) (*models.Transaction, *models.Wallet, error)
	FindBooking(bookingId uint) (*models.Reservation, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WalletConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(b *WalletConfig) WalletRepository {
	return &walletRepository{db: b.DB}
}

func (w *walletRepository) FindWalletFromUserId(id uint) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := w.db.Table("wallets").Where("user_id=?", id).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (w *walletRepository) TopUp(d *dto.TopUpRequest) (bool, error) {
	var wallet *models.Wallet
	query := w.db.Table("wallets").Where("user_id=?", d.UserID).First(&wallet)
	query = w.db.Table("wallets").Where("user_id=?", d.UserID).Update("balance", wallet.Balance+d.Amount)
	chances := 0
	chances += int(math.Floor(float64(d.Amount / 500000)))
	var gamesChances *models.GamesChances
	query = w.db.Table("games_chances").Where("user_id=?", d.UserID).First(&gamesChances)
	query = w.db.Table("games_chances").Where("user_id=?", d.UserID).Update("chance", gamesChances.Chance+chances)
	if query.Error != nil {
		return false, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED ERROR",
			Message:    "Unauthorized error.",
		}
	}
	return true, nil
}

func (w *walletRepository) CheckTransaction(bookingId uint) (bool, error) {
	var transaction *models.Transaction
	w.db.Table("transactions").Where("reservation_id=?", bookingId).First(&transaction)
	if transaction.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (w *walletRepository) FindBooking(bookingId uint) (*models.Reservation, error) {
	var res *models.Reservation
	w.db.Table("reservations").Where("id = ?", bookingId).First(&res)
	return res, nil
}

func (w *walletRepository) CountTotalTransactions() (int, error) {
	var countTransactions int64
	w.db.Table("transactions").Count(&countTransactions)
	result := int(countTransactions)
	return result, nil
}

func (w *walletRepository) CheckBalance(userId uint, price int) (bool, error) {
	var wallet *models.Wallet
	w.db.Table("wallets").Where("user_id=?", userId).First(&wallet)
	if wallet.Balance < price {
		return false, nil
	}
	return true, nil
}

func (w *walletRepository) PayBooking(userId, transactionId, bookingId uint, price int) (*models.Transaction, *models.Wallet, error) {
	var wallet *models.Wallet
	res := w.db.Table("wallets").Where("user_id=?", userId).First(&wallet)
	res = w.db.Table("wallets").Where("user_id=?", userId).Update("balance", wallet.Balance-price)
	res = w.db.Table("wallets").Where("user_id=?", userId).First(&wallet)
	if res.Error != nil {
		return nil, nil, res.Error
	}
	var reservation *models.Reservation
	res = w.db.Table("reservations").Where("id=?", bookingId).First(&reservation)
	newTransaction := models.Transaction{
		ID:            transactionId,
		HouseID:       reservation.HouseID,
		UserID:        userId,
		ReservationID: reservation.ID,
	}
	w.db.Table("transactions").Clauses(clause.OnConflict{DoNothing: true}).Create(&newTransaction)
	if res.Error != nil {
		return nil, nil, res.Error
	}
	return &newTransaction, wallet, nil
}
