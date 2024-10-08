// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hrvadl/book-service/internal/domain/recommendation (interfaces: ReadingHistorySource)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_history.go -package=mocks . ReadingHistorySource
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	history "github.com/hrvadl/book-service/internal/domain/history"
	gomock "go.uber.org/mock/gomock"
)

// MockReadingHistorySource is a mock of ReadingHistorySource interface.
type MockReadingHistorySource struct {
	ctrl     *gomock.Controller
	recorder *MockReadingHistorySourceMockRecorder
}

// MockReadingHistorySourceMockRecorder is the mock recorder for MockReadingHistorySource.
type MockReadingHistorySourceMockRecorder struct {
	mock *MockReadingHistorySource
}

// NewMockReadingHistorySource creates a new mock instance.
func NewMockReadingHistorySource(ctrl *gomock.Controller) *MockReadingHistorySource {
	mock := &MockReadingHistorySource{ctrl: ctrl}
	mock.recorder = &MockReadingHistorySourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadingHistorySource) EXPECT() *MockReadingHistorySourceMockRecorder {
	return m.recorder
}

// GetByUserID mocks base method.
func (m *MockReadingHistorySource) GetByUserID(arg0 context.Context, arg1 int) ([]history.ReadingHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", arg0, arg1)
	ret0, _ := ret[0].([]history.ReadingHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockReadingHistorySourceMockRecorder) GetByUserID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockReadingHistorySource)(nil).GetByUserID), arg0, arg1)
}
