package repositories

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type HouseRepository interface {
	DateFilter(houses []*models.House, startDate, endDate *time.Time) []*models.House
	FindHouses(searchType, search, sortBy, sortType, limit, page string, start, end *time.Time) ([]*models.House, error)
	FindHouse(id int) (*models.House, error)
	CountTotalBookings() (int, error)
	CountTotalPrice(nights int, id uint) (int, error)
	SaveBooking(booking *models.Reservation) (*models.Reservation, error)
	CheckAvailability(houseId uint, sD, eD *time.Time) (bool, error)
	FindBookings(id uint) ([]*models.Reservation, error)
	CountTotalHouses() (int, error)
	AddHouse(house *models.House) (*models.House, error)
	UpdateHouse(userId, houseId uint, input *dto.UpdateHouseRequest) (*models.House, error)
	DeleteHouse(id, houseId uint) (*models.House, error)
	GetPhotosByHouseId(houseId uint) ([]*models.HousePhoto, error)
	GetFirstPhoto(houseId uint) (*models.HousePhoto, error)
	GetBookingById(resId int) (*models.Reservation, error)
	SavePhotos(houseId uint, photos []*models.HousePhoto) ([]*models.HousePhoto, error)
	GetTotalPhotos() (int, error)
	CheckRole(userId uint) (string, error)
	GetCities() ([]*models.City, error)
	GetTransactions(userId uint) ([]*models.Transaction, error)
}

type houseRepository struct {
	db *gorm.DB
}

type HouseConfig struct {
	DB *gorm.DB
}

func NewHouseRepository(b *HouseConfig) HouseRepository {
	return &houseRepository{db: b.DB}
}

func (h *houseRepository) DateFilter(houses []*models.House, startDate, endDate *time.Time) []*models.House {
	if startDate == nil && endDate == nil {
		return houses
	}
	filteredHouses := make([]*models.House, 0)
	for _, house := range houses {
		b, _ := h.CheckAvailability(house.ID, startDate, endDate)
		if b {
			filteredHouses = append(filteredHouses, house)
		}
	}
	return filteredHouses
}

func (h *houseRepository) FindHouses(searchType, search, sortBy, sortType, limit, page string, startDate, endDate *time.Time) ([]*models.House, error) {
	houses := make([]*models.House, 0)

	defaultLimit := 8
	if limit != "" {
		defaultLimit, _ = strconv.Atoi(limit)
	}
	pageNum := 1
	if page != "" {
		pageNum, _ = strconv.Atoi(page)
	}

	offset := defaultLimit * (pageNum - 1)
	if search != "" {
		if searchType == "name" {
			if sortBy != "" {
				if sortBy == "name" {
					h.db.Table("houses").Order("house_name "+sortType).Limit(defaultLimit).Offset(offset).Where("house_name ILIKE ?", "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
				if sortBy == "price" {
					h.db.Table("houses").Order("price_per_night "+sortType).Limit(defaultLimit).Offset(offset).Where("house_name ILIKE ?", "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
				if sortBy == "city" {
					h.db.Table(`"houses"`).Order("city_name "+sortType).Limit(defaultLimit).Offset(offset).Where(`"house_name" ILIKE ?`, "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
			}
			h.db.Table("houses").Limit(defaultLimit).Offset(offset).Where("house_name ILIKE ?", "%"+search+"%").Find(&houses)
			return h.DateFilter(houses, startDate, endDate), nil
		}
		if searchType == "city" {
			if sortBy != "" {
				if sortBy == "name" {
					h.db.Table("houses").Order("house_name "+sortType).Limit(defaultLimit).Offset(offset).Where("city_name ILIKE ?", "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
				if sortBy == "price" {
					h.db.Table("houses").Order("price_per_night "+sortType).Limit(defaultLimit).Offset(offset).Where("city_name ILIKE ?", "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
				if sortBy == "city" {
					h.db.Table("houses").Order("city_name "+sortType).Limit(defaultLimit).Offset(offset).Where("city_name ILIKE ?", "%"+search+"%").Find(&houses)
					return h.DateFilter(houses, startDate, endDate), nil
				}
			}
			h.db.Table("houses").Order("created_at desc").Limit(defaultLimit).Offset(offset).Where("city_name ILIKE ?", "%"+search+"%").Find(&houses)
			return h.DateFilter(houses, startDate, endDate), nil
		}
	}
	if sortBy == "name" {
		h.db.Table("houses").Order("house_name " + sortType).Limit(defaultLimit).Offset(offset).Find(&houses)
		return h.DateFilter(houses, startDate, endDate), nil
	}
	if sortBy == "price" {
		h.db.Table("houses").Order("price_per_night " + sortType).Limit(defaultLimit).Offset(offset).Find(&houses)
		return h.DateFilter(houses, startDate, endDate), nil
	}
	if sortBy == "city" {
		h.db.Table("houses").Order("city_name " + sortType).Limit(defaultLimit).Offset(offset).Find(&houses)
		return h.DateFilter(houses, startDate, endDate), nil
	}
	h.db.Table("houses").Order("created_at desc").Limit(defaultLimit).Offset(offset).Find(&houses)
	return h.DateFilter(houses, startDate, endDate), nil
}

func (h *houseRepository) FindHouse(id int) (*models.House, error) {
	var house *models.House
	h.db.Table("houses").Where("id = ?", id).First(&house)
	return house, nil
}

func (h *houseRepository) CountTotalBookings() (int, error) {
	var totalBookings int64
	h.db.Table("reservations").Count(&totalBookings)
	return int(totalBookings), nil
}

func (h *houseRepository) CountTotalPrice(nights int, id uint) (int, error) {
	var house *models.House
	h.db.Table("houses").Where("id = ?", id).First(&house)
	return house.PricePerNight * nights, nil
}

func (h *houseRepository) SaveBooking(booking *models.Reservation) (*models.Reservation, error) {
	result := h.db.Table("reservations").Clauses(clause.OnConflict{DoNothing: true}).Create(booking)
	return booking, result.Error
}

func (h *houseRepository) CheckAvailability(houseId uint, sD, eD *time.Time) (bool, error) {
	s, _ := time.Parse("2006-01-02T15:04:05", sD.Format("2006-01-02T15:04:05"))
	e, _ := time.Parse("2006-01-02T15:04:05", eD.Format("2006-01-02T15:04:05"))
	reservations := make([]*models.Reservation, 0)
	h.db.Table("reservations").Where("house_id = ?", houseId).Find(&reservations)
	for _, reservation := range reservations {
		if inTimeSpan(reservation.CheckInDate, reservation.CheckOutDate, s, e) {
			return false, nil
		}
	}
	return true, nil
}

func inTimeSpan(start, end, checkStart, checkEnd time.Time) bool {
	return checkStart.After(start.Add(-1*time.Hour)) && checkStart.Before(end.Add(1*time.Hour)) ||
		checkEnd.After(start.Add(-1*time.Hour)) && checkEnd.Before(end.Add(1*time.Hour))
}

func (h *houseRepository) FindBookings(id uint) ([]*models.Reservation, error) {
	bookings := make([]*models.Reservation, 0)
	h.db.Table("reservations").Where("user_id = ?", id).Find(&bookings)
	return bookings, nil
}

func (h *houseRepository) CountTotalHouses() (int, error) {
	var countHouses int64
	h.db.Table("houses").Count(&countHouses)
	return int(countHouses), nil
}

func (h *houseRepository) AddHouse(house *models.House) (*models.House, error) {
	h.db.Table("houses").Create(house)
	return house, nil
}

func (h *houseRepository) UpdateHouse(userId, houseId uint, input *dto.UpdateHouseRequest) (*models.House, error) {
	var house *models.House
	h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).First(&house)
	if input.HouseName != "" {
		h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).Update("house_name", input.HouseName)
	}
	if input.PricePerNight != 0 {
		h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).Update("price_per_night", input.PricePerNight)
	}
	if input.Description != "" {
		h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).Update("description", input.Description)
	}
	if input.MaxGuest != 0 {
		h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).Update("max_guest", input.MaxGuest)
	}
	h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).First(&house)
	return house, nil
}

func (h *houseRepository) DeleteHouse(userId, houseId uint) (*models.House, error) {
	var house *models.House
	if userId == 1 {
		err := h.db.Table("houses").Where("id = ?", houseId).First(&house).Error
		h.db.Table("houses").Where("id = ?", houseId).Delete(&house)
		h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).First(&house)
		return house, err
	}
	err := h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).First(&house).Error
	err = h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).Delete(&house).Error
	h.db.Table("houses").Where("id = ?", houseId).Where("user_id = ?", userId).First(&house)
	return house, err
}

func (h *houseRepository) GetPhotosByHouseId(houseId uint) ([]*models.HousePhoto, error) {
	photos := make([]*models.HousePhoto, 0)
	q := h.db.Table("houses_photos").Where("house_id=?", houseId).Find(&photos)
	return photos, q.Error
}

func (h *houseRepository) GetFirstPhoto(houseId uint) (*models.HousePhoto, error) {
	var photo *models.HousePhoto
	q := h.db.Table("houses_photos").Where("house_id=?", houseId).First(&photo)
	if photo == nil {
		q := h.db.Table("houses_photos").Where("id=?", 550001).First(&photo)
		return photo, q.Error
	}
	return photo, q.Error
}

func (h *houseRepository) GetBookingById(resId int) (*models.Reservation, error) {
	var res *models.Reservation
	q := h.db.Table("reservations").Where("id=?", resId).First(&res)
	return res, q.Error
}

func (h *houseRepository) SavePhotos(houseId uint, photos []*models.HousePhoto) ([]*models.HousePhoto, error) {
	var responsePhotos []*models.HousePhoto
	for _, photo := range photos {
		s := h.db.Table("houses_photos").Where("house_id=?", houseId).Create(&photo)
		if s.Error != nil {
			return nil, s.Error
		}
		responsePhotos = append(responsePhotos, photo)
	}
	return responsePhotos, nil
}

func (h *houseRepository) GetTotalPhotos() (int, error) {
	var totalPhotos int64
	q := h.db.Table("houses_photos").Count(&totalPhotos)
	return int(totalPhotos), q.Error
}

func (h *houseRepository) CheckRole(userId uint) (string, error) {
	var user *models.User
	q := h.db.Table("users").Where("id=?", userId).First(&user)
	return user.Role, q.Error
}

func (h *houseRepository) GetCities() ([]*models.City, error) {
	cities := make([]*models.City, 0)
	q := h.db.Table("cities").Find(&cities)
	if q.Error != nil {
		return nil, q.Error
	}
	return cities, nil
}

func (h *houseRepository) GetTransactions(userId uint) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	q := h.db.Table("transactions").Where("user_id", userId).Find(&transactions)
	if q.Error != nil {
		return nil, q.Error
	}
	return transactions, nil
}
