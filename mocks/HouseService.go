// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"

	mock "github.com/stretchr/testify/mock"
)

// HouseService is an autogenerated mock type for the HouseService type
type HouseService struct {
	mock.Mock
}

// AddHouse provides a mock function with given fields: userId, input
func (_m *HouseService) AddHouse(userId uint, input *dto.AddHouseRequest) (*dto.AddHouseResponse, error) {
	ret := _m.Called(userId, input)

	var r0 *dto.AddHouseResponse
	if rf, ok := ret.Get(0).(func(uint, *dto.AddHouseRequest) *dto.AddHouseResponse); ok {
		r0 = rf(userId, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.AddHouseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *dto.AddHouseRequest) error); ok {
		r1 = rf(userId, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckRole provides a mock function with given fields: userId
func (_m *HouseService) CheckRole(userId uint) (string, error) {
	ret := _m.Called(userId)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReservation provides a mock function with given fields: houseId, userId, request
func (_m *HouseService) CreateReservation(houseId uint, userId uint, request *dto.BookHouseRequest) (*dto.BookingsResponse, error) {
	ret := _m.Called(houseId, userId, request)

	var r0 *dto.BookingsResponse
	if rf, ok := ret.Get(0).(func(uint, uint, *dto.BookHouseRequest) *dto.BookingsResponse); ok {
		r0 = rf(houseId, userId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.BookingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, *dto.BookHouseRequest) error); ok {
		r1 = rf(houseId, userId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHouse provides a mock function with given fields: userId, houseId
func (_m *HouseService) DeleteHouse(userId uint, houseId uint) (*dto.DeleteHouseResponse, error) {
	ret := _m.Called(userId, houseId)

	var r0 *dto.DeleteHouseResponse
	if rf, ok := ret.Get(0).(func(uint, uint) *dto.DeleteHouseResponse); ok {
		r0 = rf(userId, houseId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.DeleteHouseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userId, houseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateBookingsId provides a mock function with given fields:
func (_m *HouseService) GenerateBookingsId() (uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateHouseId provides a mock function with given fields:
func (_m *HouseService) GenerateHouseId() (uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookingById provides a mock function with given fields: resId
func (_m *HouseService) GetBookingById(resId int) (*dto.GetBookingResponse, error) {
	ret := _m.Called(resId)

	var r0 *dto.GetBookingResponse
	if rf, ok := ret.Get(0).(func(int) *dto.GetBookingResponse); ok {
		r0 = rf(resId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetBookingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(resId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookings provides a mock function with given fields: id
func (_m *HouseService) GetBookings(id uint) (*dto.ShowBookingsResponse, error) {
	ret := _m.Called(id)

	var r0 *dto.ShowBookingsResponse
	if rf, ok := ret.Get(0).(func(uint) *dto.ShowBookingsResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ShowBookingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFirstPhoto provides a mock function with given fields: u
func (_m *HouseService) GetFirstPhoto(u uint) (*dto.FirstPhotoResponse, error) {
	ret := _m.Called(u)

	var r0 *dto.FirstPhotoResponse
	if rf, ok := ret.Get(0).(func(uint) *dto.FirstPhotoResponse); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.FirstPhotoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHouse provides a mock function with given fields: id
func (_m *HouseService) GetHouse(id int) (*dto.GetHouseResponse, error) {
	ret := _m.Called(id)

	var r0 *dto.GetHouseResponse
	if rf, ok := ret.Get(0).(func(int) *dto.GetHouseResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetHouseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHouses provides a mock function with given fields: request, searchType, search, sortBy, sortType, limit, page
func (_m *HouseService) ShowHouses(request *dto.BookHouseRequest, searchType string, search string, sortBy string, sortType string, limit string, page string) (*dto.HousesResponse, error) {
	ret := _m.Called(request, searchType, search, sortBy, sortType, limit, page)

	var r0 *dto.HousesResponse
	if rf, ok := ret.Get(0).(func(*dto.BookHouseRequest, string, string, string, string, string, string) *dto.HousesResponse); ok {
		r0 = rf(request, searchType, search, sortBy, sortType, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.HousesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.BookHouseRequest, string, string, string, string, string, string) error); ok {
		r1 = rf(request, searchType, search, sortBy, sortType, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPhotosByHouseId provides a mock function with given fields: u
func (_m *HouseService) GetPhotosByHouseId(u uint) (*dto.PhotosResponse, error) {
	ret := _m.Called(u)

	var r0 *dto.PhotosResponse
	if rf, ok := ret.Get(0).(func(uint) *dto.PhotosResponse); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PhotosResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHouse provides a mock function with given fields: userId, houseId, input
func (_m *HouseService) UpdateHouse(userId uint, houseId uint, input *dto.UpdateHouseRequest) (*dto.UpdateHouseResponse, error) {
	ret := _m.Called(userId, houseId, input)

	var r0 *dto.UpdateHouseResponse
	if rf, ok := ret.Get(0).(func(uint, uint, *dto.UpdateHouseRequest) *dto.UpdateHouseResponse); ok {
		r0 = rf(userId, houseId, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateHouseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, *dto.UpdateHouseRequest) error); ok {
		r1 = rf(userId, houseId, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPhotos provides a mock function with given fields: houseId, input
func (_m *HouseService) UploadPhotos(houseId uint, input *dto.UploadPhotosRequest) (*dto.UploadPhotosResponse, error) {
	ret := _m.Called(houseId, input)

	var r0 *dto.UploadPhotosResponse
	if rf, ok := ret.Get(0).(func(uint, *dto.UploadPhotosRequest) *dto.UploadPhotosResponse); ok {
		r0 = rf(houseId, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UploadPhotosResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *dto.UploadPhotosRequest) error); ok {
		r1 = rf(houseId, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHouseService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHouseService creates a new instance of HouseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHouseService(t mockConstructorTestingTNewHouseService) *HouseService {
	mock := &HouseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
