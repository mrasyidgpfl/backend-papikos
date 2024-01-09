// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"

	mock "github.com/stretchr/testify/mock"

	models "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// CheckBalance provides a mock function with given fields: userId, price
func (_m *WalletRepository) CheckBalance(userId uint, price int) (bool, error) {
	ret := _m.Called(userId, price)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint, int) bool); ok {
		r0 = rf(userId, price)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, int) error); ok {
		r1 = rf(userId, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTransaction provides a mock function with given fields: bookingId
func (_m *WalletRepository) CheckTransaction(bookingId uint) (bool, error) {
	ret := _m.Called(bookingId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint) bool); ok {
		r0 = rf(bookingId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(bookingId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountTotalTransactions provides a mock function with given fields:
func (_m *WalletRepository) CountTotalTransactions() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBooking provides a mock function with given fields: bookingId
func (_m *WalletRepository) FindBooking(bookingId uint) (*models.Reservation, error) {
	ret := _m.Called(bookingId)

	var r0 *models.Reservation
	if rf, ok := ret.Get(0).(func(uint) *models.Reservation); ok {
		r0 = rf(bookingId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(bookingId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWalletFromUserId provides a mock function with given fields: id
func (_m *WalletRepository) FindWalletFromUserId(id uint) (*models.Wallet, error) {
	ret := _m.Called(id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(uint) *models.Wallet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
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

// PayBooking provides a mock function with given fields: userId, transactionId, bookingId, price
func (_m *WalletRepository) PayBooking(userId uint, transactionId uint, bookingId uint, price int) (*models.Transaction, *models.Wallet, error) {
	ret := _m.Called(userId, transactionId, bookingId, price)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(uint, uint, uint, int) *models.Transaction); ok {
		r0 = rf(userId, transactionId, bookingId, price)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 *models.Wallet
	if rf, ok := ret.Get(1).(func(uint, uint, uint, int) *models.Wallet); ok {
		r1 = rf(userId, transactionId, bookingId, price)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Wallet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint, uint, uint, int) error); ok {
		r2 = rf(userId, transactionId, bookingId, price)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TopUp provides a mock function with given fields: d
func (_m *WalletRepository) TopUp(d *dto.TopUpRequest) (bool, error) {
	ret := _m.Called(d)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*dto.TopUpRequest) bool); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.TopUpRequest) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}