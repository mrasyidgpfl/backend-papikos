// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"

	mock "github.com/stretchr/testify/mock"

	models "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// BlackListToken provides a mock function with given fields: token
func (_m *AuthRepository) BlackListToken(token string) (bool, error) {
	ret := _m.Called(token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckToken provides a mock function with given fields: token
func (_m *AuthRepository) CheckToken(token string) (bool, error) {
	ret := _m.Called(token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanBlacklists provides a mock function with given fields:
func (_m *AuthRepository) CleanBlacklists() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountTotalUser provides a mock function with given fields:
func (_m *AuthRepository) CountTotalUser() (int, error) {
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

// FindUserDetailsById provides a mock function with given fields: userId
func (_m *AuthRepository) FindUserDetailsById(userId uint) (*models.User, *models.Wallet, *models.GamesChances, error) {
	ret := _m.Called(userId)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(uint) *models.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 *models.Wallet
	if rf, ok := ret.Get(1).(func(uint) *models.Wallet); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Wallet)
		}
	}

	var r2 *models.GamesChances
	if rf, ok := ret.Get(2).(func(uint) *models.GamesChances); ok {
		r2 = rf(userId)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*models.GamesChances)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(uint) error); ok {
		r3 = rf(userId)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MatchingCredential provides a mock function with given fields: email, password
func (_m *AuthRepository) MatchingCredential(email string, password string) (*models.User, error) {
	ret := _m.Called(email, password)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string, string) *models.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveGamesChances provides a mock function with given fields: gamesChances
func (_m *AuthRepository) SaveGamesChances(gamesChances *models.GamesChances) (*models.GamesChances, int, error) {
	ret := _m.Called(gamesChances)

	var r0 *models.GamesChances
	if rf, ok := ret.Get(0).(func(*models.GamesChances) *models.GamesChances); ok {
		r0 = rf(gamesChances)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GamesChances)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*models.GamesChances) int); ok {
		r1 = rf(gamesChances)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.GamesChances) error); ok {
		r2 = rf(gamesChances)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveUser provides a mock function with given fields: user
func (_m *AuthRepository) SaveUser(user *models.User) (*models.User, int, error) {
	ret := _m.Called(user)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*models.User) *models.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*models.User) int); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.User) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveWallet provides a mock function with given fields: waller
func (_m *AuthRepository) SaveWallet(waller *models.Wallet) (*models.Wallet, int, error) {
	ret := _m.Called(waller)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(*models.Wallet) *models.Wallet); ok {
		r0 = rf(waller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*models.Wallet) int); ok {
		r1 = rf(waller)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.Wallet) error); ok {
		r2 = rf(waller)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateProfile provides a mock function with given fields: request
func (_m *AuthRepository) UpdateProfile(request *dto.UpdateProfileRequest) (*models.User, error) {
	ret := _m.Called(request)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*dto.UpdateProfileRequest) *models.User); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.UpdateProfileRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository(t mockConstructorTestingTNewAuthRepository) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
