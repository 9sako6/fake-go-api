// Code generated by MockGen. DO NOT EDIT.
// Source: driver/user.go
//
// Generated by this command:
//
//	mockgen -source=driver/user.go -destination=mock/mock_driver/mock_user_driver.go
//
// Package mock_driver is a generated GoMock package.
package mock_driver

import (
	driver "fake_bff/driver"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserDvier is a mock of UserDvier interface.
type MockUserDvier struct {
	ctrl     *gomock.Controller
	recorder *MockUserDvierMockRecorder
}

// MockUserDvierMockRecorder is the mock recorder for MockUserDvier.
type MockUserDvierMockRecorder struct {
	mock *MockUserDvier
}

// NewMockUserDvier creates a new mock instance.
func NewMockUserDvier(ctrl *gomock.Controller) *MockUserDvier {
	mock := &MockUserDvier{ctrl: ctrl}
	mock.recorder = &MockUserDvierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDvier) EXPECT() *MockUserDvierMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserDvier) GetUserById(userId string) (*driver.GetUserByIdQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", userId)
	ret0, _ := ret[0].(*driver.GetUserByIdQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserDvierMockRecorder) GetUserById(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserDvier)(nil).GetUserById), userId)
}