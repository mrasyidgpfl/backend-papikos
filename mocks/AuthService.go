// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"

	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// CheckToken provides a mock function with given fields: token
func (_m *AuthService) CheckToken(token string) (bool, error) {
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
func (_m *AuthService) CleanBlacklists() (*dto.CleanBlacklistResponse, error) {
	ret := _m.Called()

	var r0 *dto.CleanBlacklistResponse
	if rf, ok := ret.Get(0).(func() *dto.CleanBlacklistResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CleanBlacklistResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDetail provides a mock function with given fields: id
func (_m *AuthService) GetUserDetail(id uint) (*dto.UserDetailResponse, error) {
	ret := _m.Called(id)

	var r0 *dto.UserDetailResponse
	if rf, ok := ret.Get(0).(func(uint) *dto.UserDetailResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDetailResponse)
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

// SignIn provides a mock function with given fields: _a0
func (_m *AuthService) SignIn(_a0 *dto.SignInRequest) (*dto.SignInResponse, error) {
	ret := _m.Called(_a0)

	var r0 *dto.SignInResponse
	if rf, ok := ret.Get(0).(func(*dto.SignInRequest) *dto.SignInResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SignInResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.SignInRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignOut provides a mock function with given fields: token
func (_m *AuthService) SignOut(token string) (*dto.SignOutResponse, error) {
	ret := _m.Called(token)

	var r0 *dto.SignOutResponse
	if rf, ok := ret.Get(0).(func(string) *dto.SignOutResponse); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SignOutResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: d
func (_m *AuthService) SignUp(d *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	ret := _m.Called(d)

	var r0 *dto.SignUpResponse
	if rf, ok := ret.Get(0).(func(*dto.SignUpRequest) *dto.SignUpResponse); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SignUpResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.SignUpRequest) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: d
func (_m *AuthService) UpdateProfile(d *dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error) {
	ret := _m.Called(d)

	var r0 *dto.UpdateProfileResponse
	if rf, ok := ret.Get(0).(func(*dto.UpdateProfileRequest) *dto.UpdateProfileResponse); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateProfileResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.UpdateProfileRequest) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthService(t mockConstructorTestingTNewAuthService) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}