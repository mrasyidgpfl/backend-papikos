package services

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"fmt"
	"net/http"
)

type PickUpService interface {
	PickUp(id uint, d *dto.PickUpRequest) (*dto.PickUpResponse, error)
	GeneratePickUpIds() (uint, uint, error)
	UpdatePUStatus(request *dto.UpdatePickUpStatusRequest) (*dto.UpdatePickUpStatusResponse, error)
	GetPickUps(userId uint) (*dto.GetPickUps, error)
}

type pickUpService struct {
	pickUpRepository repositories.PickUpRepository
	appConfig        config.AppConfig
}

type PickUpConfig struct {
	PickUpRepository repositories.PickUpRepository
	AppConfig        config.AppConfig
}

func NewPickUpService(c *PickUpConfig) PickUpService {
	return &pickUpService{
		pickUpRepository: c.PickUpRepository,
		appConfig:        c.AppConfig,
	}
}

func (p *pickUpService) PickUp(userId uint, request *dto.PickUpRequest) (*dto.PickUpResponse, error) {
	pickUpId, pSId, _ := p.GeneratePickUpIds()
	status := models.PickUpStatus{ID: pSId, Status: "Pending Admin"}
	price, _ := p.pickUpRepository.CalculatePickUpPrice(userId, request.ReservationID)
	if price == 0 {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "ALREADY PAID ERROR",
			Message:    "Pick-up service has already been booked and paid",
		}
	}
	checkBalance, wallet, _ := p.pickUpRepository.CheckBalance(userId, price)
	if !checkBalance {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "INSUFFICIENT BALANCE ERROR",
			Message:    "Balance insufficient, please top up",
		}
	}
	_, paymentErr := p.pickUpRepository.PayForPickUp(userId, price)
	if paymentErr != nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusInternalServerError,
			Code:       "PAYMENT UNSUCCESSFUL ERROR",
			Message:    "Payment unsuccessful, please try again",
		}
	}
	_, _ = p.pickUpRepository.AddStatus(&status)
	pickUp, _ := p.pickUpRepository.CreatePickUp(pickUpId, userId, request, pSId)
	return new(dto.PickUpResponse).CreatePickUpResponse(200, "PICK-UP BOOKING SUCCESSFUL", "Pick up booking successfully created and already paid, status is now Pending Admin", pickUp, wallet), nil
}

func (p *pickUpService) GeneratePickUpIds() (uint, uint, error) {
	totalPickUp, _ := p.pickUpRepository.CountTotalPickUps()
	totalPickUp += 1
	pickUpId := 700000 + totalPickUp
	pSId := 750000 + totalPickUp
	return uint(pickUpId), uint(pSId), nil
}

func (p *pickUpService) UpdatePUStatus(request *dto.UpdatePickUpStatusRequest) (*dto.UpdatePickUpStatusResponse, error) {
	pUStatus, err := p.pickUpRepository.UpdateStatus(request)
	return new(dto.UpdatePickUpStatusResponse).CreateUpdatePickUpStatusRes(200, "STATUS UPDATE SUCCESSFUL", "Status successfully updated", pUStatus), err
}

func (p *pickUpService) GetPickUps(userId uint) (*dto.GetPickUps, error) {
	pickups, err := p.pickUpRepository.GetPickups(userId)
	if err != nil {
		return nil, err
	}
	res := make([]*dto.PURes, 0)
	for _, pickup := range pickups {
		fmt.Println(pickup)
		status, err := p.pickUpRepository.GetPUStatus(pickup.PickUpStatusID)
		if err != nil {
			return nil, err
		}
		p := dto.PURes{
			PickUp: pickup,
			Status: status,
		}
		res = append(res, &p)
	}
	return new(dto.GetPickUps).GetPickUps(200, "PICKUPS RETRIEVAL SUCCESSFUL", "Pickups successfully retrieved", res), err
}
