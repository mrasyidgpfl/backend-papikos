package services

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"math"
	"net/http"
	"time"
)

type HouseService interface {
	GenerateBookingsId() (uint, error)
	GetHouses(searchType, search, sortBy, sortType, start, end, limit, page string) (*dto.HousesResponse, error)
	GetHouse(id int) (*dto.GetHouseResponse, error)
	CreateReservation(houseId uint, userId uint, request *dto.BookHouseRequest) (*dto.BookingsResponse, error)
	GetBookings(id uint) (*dto.ShowBookingsResponse, error)
	AddHouse(userId uint, input *dto.AddHouseRequest) (*dto.AddHouseResponse, error)
	GenerateHouseId() (uint, error)
	UpdateHouse(userId, houseId uint, input *dto.UpdateHouseRequest) (*dto.UpdateHouseResponse, error)
	DeleteHouse(userId, houseId uint) (*dto.DeleteHouseResponse, error)
	GetPhotosByHouseId(u uint) (*dto.PhotosResponse, error)
	GetFirstPhoto(u uint) (*dto.FirstPhotoResponse, error)
	GetBookingById(resId int) (*dto.GetBookingResponse, error)
	UploadPhotos(houseId uint, input *dto.UploadPhotosRequest) (*dto.UploadPhotosResponse, error)
	CheckRole(userId uint) (string, error)
	GetCities() (*dto.GetCitiesResponse, error)
	GetTransactions(userId uint) (*dto.GetTransactions, error)
}

type houseService struct {
	houseRepository repositories.HouseRepository
	appConfig       config.AppConfig
}

type HouseConfig struct {
	HouseRepository repositories.HouseRepository
	AppConfig       config.AppConfig
}

func NewHouseService(c *HouseConfig) HouseService {
	return &houseService{
		houseRepository: c.HouseRepository,
		appConfig:       c.AppConfig,
	}
}

func (h *houseService) GetHouses(searchType, search, sortBy, sortType, start, end, limit, page string) (*dto.HousesResponse, error) {
	hs := make([]*dto.House, 0)
	startDate, _ := time.Parse("2006-01-02T15:04:05", start)
	endDate, _ := time.Parse("2006-01-02T15:04:05", end)

	houses, err := h.houseRepository.FindHouses(searchType, search, sortBy, sortType, limit, page, &startDate, &endDate)
	for _, house := range houses {
		hs = append(hs, new(dto.House).CreateHouseResponse(house))
	}
	if start == "" && end == "" {
		start := time.Now()
		return new(dto.House).ShowHouses(200, "HOUSES FOUND", "Houses successfully retrieved", hs, &start, nil), err
	}
	return new(dto.House).ShowHouses(200, "HOUSES FOUND", "Houses successfully retrieved", hs, &startDate, &endDate), err
}

func (h *houseService) GetHouse(id int) (*dto.GetHouseResponse, error) {
	house, err := h.houseRepository.FindHouse(id)
	return new(dto.GetHouseResponse).CreateGetHouseResponse(200, "HOUSE RETRIEVAL SUCCESS", "House retrieved successfully", house), err
}

func (h *houseService) GenerateBookingsId() (uint, error) {
	totalUser, _ := h.houseRepository.CountTotalBookings()
	totalUser += 1
	bookingId := 400000 + totalUser
	return uint(bookingId), nil
}

func (h *houseService) CreateReservation(houseId uint, userId uint, request *dto.BookHouseRequest) (*dto.BookingsResponse, error) {
	availability, _ := h.houseRepository.CheckAvailability(houseId, &request.CheckInDate, &request.CheckOutDate)
	if availability == false {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "UNAVAILABLE LOT ERROR",
			Message:    "House already reserved, choose another date",
		}
	}
	bookingId, _ := h.GenerateBookingsId()
	nights := int(math.Ceil(request.CheckOutDate.Sub(request.CheckInDate).Hours() / 24))
	totalPrice, _ := h.houseRepository.CountTotalPrice(nights, houseId)
	newBooking := models.Reservation{
		ID:           bookingId,
		HouseID:      houseId,
		UserID:       userId,
		CheckInDate:  request.CheckInDate,
		CheckOutDate: request.CheckOutDate,
		TotalPrice:   totalPrice,
	}
	reservation, _ := h.houseRepository.SaveBooking(&newBooking)
	if reservation == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "BAD REQUEST ERROR",
			Message:    "Bad request error",
		}
	}
	return new(dto.BookingsResponse).CreateBookingResponse(200, "HOUSE BOOKING SUCCESS", "House booked successfully", reservation), nil
}

func (h *houseService) GetBookings(id uint) (*dto.ShowBookingsResponse, error) {
	bookings, err := h.houseRepository.FindBookings(id)
	return new(dto.ShowBookingsResponse).ShowBookings(200, "BOOKINGS FOUND", "Bookings successfully retrieved", bookings), err
}

func (h *houseService) AddHouse(userId uint, input *dto.AddHouseRequest) (*dto.AddHouseResponse, error) {
	houseId, _ := h.GenerateHouseId()
	house := models.House{
		ID:            houseId,
		HouseName:     input.HouseName,
		UserID:        userId,
		PricePerNight: input.PricePerNight,
		Description:   input.Description,
		CityID:        input.CityID,
		CityName:      input.CityName,
		MaxGuest:      input.MaxGuest,
		CreatedAt:     time.Now(),
	}
	insertedHouse, err := h.houseRepository.AddHouse(&house)
	if err != nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "BAD REQUEST ERROR",
			Message:    "Bad request error",
		}
	}
	return new(dto.AddHouseResponse).CreateAddHouseResponse(200, "ADD HOUSE SUCCESSFUL", "House created successfully", insertedHouse), err
}

func (h *houseService) GenerateHouseId() (uint, error) {
	totalHouses, _ := h.houseRepository.CountTotalHouses()
	totalHouses += 1
	houseId := uint(500000 + totalHouses)
	return houseId, nil
}

func (h *houseService) UpdateHouse(userId, houseId uint, input *dto.UpdateHouseRequest) (*dto.UpdateHouseResponse, error) {
	updatedHouse, err := h.houseRepository.UpdateHouse(userId, houseId, input)
	if err != nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusBadRequest,
			Code:       "BAD REQUEST ERROR",
			Message:    "Bad request error.",
		}
	}
	t := time.Now()
	return new(dto.UpdateHouseResponse).CreateUpdateHouseResponse(200, "UPDATE HOUSE SUCCESSFUL", "House's information updated successfully", updatedHouse, t), err
}

func (h *houseService) DeleteHouse(userId, houseId uint) (*dto.DeleteHouseResponse, error) {
	_, err := h.houseRepository.DeleteHouse(userId, houseId)
	return new(dto.DeleteHouseResponse).CreateDeleteHouseResponse(400, "DELETE HOUSE SUCCESSFUL", "Successfully deleted house"), err
}

func (h *houseService) GetPhotosByHouseId(houseId uint) (*dto.PhotosResponse, error) {
	photos, err := h.houseRepository.GetPhotosByHouseId(houseId)
	return new(dto.PhotosResponse).CreatePhotosResponse(200, "PHOTOS RETRIEVAL SUCCESSFUL", "Photos retrieved successfully", photos), err
}

func (h *houseService) GetFirstPhoto(houseId uint) (*dto.FirstPhotoResponse, error) {
	photo, err := h.houseRepository.GetFirstPhoto(houseId)
	return new(dto.FirstPhotoResponse).CreateFirstPhotoResponse(200, "PHOTOS RETRIEVAL SUCCESSFUL", "Photos retrieved successfully", photo), err
}

func (h *houseService) GetBookingById(resId int) (*dto.GetBookingResponse, error) {
	booking, err := h.houseRepository.GetBookingById(resId)
	return new(dto.GetBookingResponse).CreateBookingResponse(200, "BOOKING RETRIEVAL SUCCESSFUL", "Booking retrieved successfully", booking), err
}

func (h *houseService) UploadPhotos(houseId uint, input *dto.UploadPhotosRequest) (*dto.UploadPhotosResponse, error) {
	var photos []*models.HousePhoto
	totalPhotos, _ := h.houseRepository.GetTotalPhotos()
	for _, url := range input.PhotosURL {
		totalPhotos += 1
		photo := models.HousePhoto{
			ID:       uint(totalPhotos),
			HouseID:  houseId,
			PhotoURL: url,
		}
		photos = append(photos, &photo)
	}
	uploadedPhotos, err := h.houseRepository.SavePhotos(houseId, photos)
	return new(dto.UploadPhotosResponse).CreateUploadPhotosResponse(200, "UPLOAD PHOTO(S) SUCCESSFUL", "Photo(s) uploaded successfully", uploadedPhotos), err
}

func (h *houseService) CheckRole(userId uint) (string, error) {
	role, err := h.houseRepository.CheckRole(userId)
	if err != nil {
		return "", err
	}
	if role == "" {
		return "", httperror.AppError{
			StatusCode: http.StatusNotFound,
			Code:       "NOT FOUND ERROR",
			Message:    "record not found",
		}
	}
	return role, nil
}

func (h *houseService) GetCities() (*dto.GetCitiesResponse, error) {
	cities, err := h.houseRepository.GetCities()
	if err != nil {
		return nil, err
	}
	return new(dto.GetCitiesResponse).CreateCitiesResponse(200, "CITIES RETRIEVAL SUCCESSFUL", "Cities successfully retrieved", cities), err
}

func (h *houseService) GetTransactions(userId uint) (*dto.GetTransactions, error) {
	transactions, err := h.houseRepository.GetTransactions(userId)
	if err != nil {
		return nil, err
	}
	return new(dto.GetTransactions).CreateTransactionsResponse(200, "TRANSACTIONS RETRIEVAL SUCCESSFUL", "Transactions successfully retrieved", transactions), err
}
