// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hrvadl/book-service/internal/domain/recommendation (interfaces: UserSource)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_users.go -package=mocks . UserSource
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	user "github.com/hrvadl/book-service/internal/domain/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserSource is a mock of UserSource interface.
type MockUserSource struct {
	ctrl     *gomock.Controller
	recorder *MockUserSourceMockRecorder
}

// MockUserSourceMockRecorder is the mock recorder for MockUserSource.
type MockUserSourceMockRecorder struct {
	mock *MockUserSource
}

// NewMockUserSource creates a new mock instance.
func NewMockUserSource(ctrl *gomock.Controller) *MockUserSource {
	mock := &MockUserSource{ctrl: ctrl}
	mock.recorder = &MockUserSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSource) EXPECT() *MockUserSourceMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockUserSource) GetByID(arg0 context.Context, arg1 int) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserSourceMockRecorder) GetByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserSource)(nil).GetByID), arg0, arg1)
}
