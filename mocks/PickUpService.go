// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"

	mock "github.com/stretchr/testify/mock"
)

// PickUpService is an autogenerated mock type for the PickUpService type
type PickUpService struct {
	mock.Mock
}

// GeneratePickUpIds provides a mock function with given fields:
func (_m *PickUpService) GeneratePickUpIds() (uint, uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 uint
	if rf, ok := ret.Get(1).(func() uint); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(uint)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PickUp provides a mock function with given fields: id, d
func (_m *PickUpService) PickUp(id uint, d *dto.PickUpRequest) (*dto.PickUpResponse, error) {
	ret := _m.Called(id, d)

	var r0 *dto.PickUpResponse
	if rf, ok := ret.Get(0).(func(uint, *dto.PickUpRequest) *dto.PickUpResponse); ok {
		r0 = rf(id, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PickUpResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *dto.PickUpRequest) error); ok {
		r1 = rf(id, d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePUStatus provides a mock function with given fields: request
func (_m *PickUpService) UpdatePUStatus(request *dto.UpdatePickUpStatusRequest) (*dto.UpdatePickUpStatusResponse, error) {
	ret := _m.Called(request)

	var r0 *dto.UpdatePickUpStatusResponse
	if rf, ok := ret.Get(0).(func(*dto.UpdatePickUpStatusRequest) *dto.UpdatePickUpStatusResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdatePickUpStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.UpdatePickUpStatusRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPickUpService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPickUpService creates a new instance of PickUpService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPickUpService(t mockConstructorTestingTNewPickUpService) *PickUpService {
	mock := &PickUpService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
