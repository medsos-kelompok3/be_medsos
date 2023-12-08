// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	user "be_medsos/features/user"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: input
func (_m *Repository) AddUser(input user.User) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userID
func (_m *Repository) DeleteUser(userID uint) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: userID
func (_m *Repository) GetUserByID(userID uint) (*user.User, error) {
	ret := _m.Called(userID)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*user.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *user.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *Repository) GetUserByUsername(username string) (user.User, error) {
	ret := _m.Called(username)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: username
func (_m *Repository) Login(username string) (user.User, error) {
	ret := _m.Called(username)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: input
func (_m *Repository) UpdateUser(input user.User) (user.User, error) {
	ret := _m.Called(input)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(user.User) (user.User, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(user.User) user.User); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(user.User) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
