package repositories

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type PickUpRepository interface {
	CountTotalPickUps() (int, error)
	AddStatus(status *models.PickUpStatus) (*models.PickUpStatus, error)
	CreatePickUp(pickUpId uint, userId uint, request *dto.PickUpRequest, pSId uint) (*models.PickUp, error)
	CalculatePickUpPrice(uid uint, r uint) (int, error)
	CheckBalance(userId uint, price int) (bool, *models.Wallet, error)
	PayForPickUp(id uint, price int) (bool, error)
	UpdateStatus(request *dto.UpdatePickUpStatusRequest) (*models.PickUpStatus, error)
	GetPickups(userId uint) ([]*models.PickUp, error)
	GetPUStatus(id uint) (string, error)
}

type pickUpRepository struct {
	db *gorm.DB
}

type PickUpConfig struct {
	DB *gorm.DB
}

func NewPickUpRepository(b *PickUpConfig) PickUpRepository {
	return &pickUpRepository{db: b.DB}
}

func (p *pickUpRepository) CountTotalPickUps() (int, error) {
	var totalPickUps int64
	p.db.Table("pickups").Count(&totalPickUps)
	return int(totalPickUps), nil
}

func (p *pickUpRepository) AddStatus(status *models.PickUpStatus) (*models.PickUpStatus, error) {
	p.db.Table("pickup_statuses").Clauses(clause.OnConflict{DoNothing: true}).Create(status)
	return status, nil
}

func (p *pickUpRepository) CreatePickUp(pickUpId uint, userId uint, request *dto.PickUpRequest, pSId uint) (*models.PickUp, error) {
	pickUp := models.PickUp{
		ID:             pickUpId,
		UserID:         userId,
		ReservationID:  request.ReservationID,
		PickUpStatusID: pSId,
	}
	p.db.Table("pickups").Clauses(clause.OnConflict{DoNothing: true}).Create(&pickUp)
	return &pickUp, nil
}

func (p *pickUpRepository) CalculatePickUpPrice(uid uint, r uint) (int, error) {
	var res *models.Reservation
	p.db.Table("reservations").Where("id=?", r).First(&res)
	var pickup *models.PickUp
	p.db.Table("pickups").Where("reservation_id=?", res.ID).First(&pickup)
	if pickup.ID != 0 {
		return 0, nil
	}
	var house *models.House
	p.db.Table("houses").Where("id=?", res.HouseID).First(&house)
	var user *models.User
	p.db.Table("users").Where("id=?", uid).First(&user)
	price := 0
	if user.CityId != house.CityID {
		price += 300000
		return price, nil
	}
	price += 100000
	return price, nil
}

func (p *pickUpRepository) CheckBalance(userId uint, price int) (bool, *models.Wallet, error) {
	var wallet *models.Wallet
	p.db.Table("wallets").Where("user_id=?", userId).First(&wallet)
	if wallet.Balance < price {
		return false, nil, nil
	}
	return true, wallet, nil
}

func (p *pickUpRepository) PayForPickUp(id uint, price int) (bool, error) {
	var wallet *models.Wallet
	res := p.db.Table("wallets").Where("user_id=?", id).First(&wallet)
	res = p.db.Table("wallets").Where("user_id=?", id).Update("balance", wallet.Balance-price)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (p *pickUpRepository) UpdateStatus(request *dto.UpdatePickUpStatusRequest) (*models.PickUpStatus, error) {
	var pU *models.PickUp
	err := p.db.Table("pickups").Where("id=?", request.PickUpID).First(&pU).Error
	pUSId := pU.PickUpStatusID
	if pUSId == 0 {
		return nil, httperror.AppError{
			StatusCode: http.StatusNotFound,
			Code:       "SOURCE NOT FOUND",
			Message:    err.Error(),
		}
	}
	var status *models.PickUpStatus
	err = p.db.Table("pickup_statuses").Where("id=?", pUSId).First(&status).Error
	err = p.db.Table("pickup_statuses").Where("id=?", pUSId).Update("status", request.Status).Error
	err = p.db.Table("pickup_statuses").Where("id=?", pUSId).First(&status).Error
	if status.Status == "" {
		return nil, httperror.AppError{
			StatusCode: http.StatusNotFound,
			Code:       "SOURCE NOT FOUND",
			Message:    err.Error(),
		}
	}
	return status, err
}

func (p *pickUpRepository) GetPickups(userId uint) ([]*models.PickUp, error) {
	pickups := make([]*models.PickUp, 0)
	q := p.db.Table("pickups").Where("user_id", userId).Find(&pickups)
	if q.Error != nil {
		return nil, q.Error
	}
	return pickups, nil
}

func (p *pickUpRepository) GetPUStatus(id uint) (string, error) {
	var status *models.PickUpStatus
	q := p.db.Table("pickup_statuses").Where("id", id).First(&status)
	if q.Error != nil {
		return "", q.Error
	}
	return status.Status, nil
}
