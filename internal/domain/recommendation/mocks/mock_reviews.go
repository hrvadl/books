// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hrvadl/book-service/internal/domain/recommendation (interfaces: ReviewSource)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_reviews.go -package=mocks . ReviewSource
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	review "github.com/hrvadl/book-service/internal/domain/review"
	gomock "go.uber.org/mock/gomock"
)

// MockReviewSource is a mock of ReviewSource interface.
type MockReviewSource struct {
	ctrl     *gomock.Controller
	recorder *MockReviewSourceMockRecorder
}

// MockReviewSourceMockRecorder is the mock recorder for MockReviewSource.
type MockReviewSourceMockRecorder struct {
	mock *MockReviewSource
}

// NewMockReviewSource creates a new mock instance.
func NewMockReviewSource(ctrl *gomock.Controller) *MockReviewSource {
	mock := &MockReviewSource{ctrl: ctrl}
	mock.recorder = &MockReviewSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewSource) EXPECT() *MockReviewSourceMockRecorder {
	return m.recorder
}

// GetByUserID mocks base method.
func (m *MockReviewSource) GetByUserID(arg0 context.Context, arg1 int) ([]review.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", arg0, arg1)
	ret0, _ := ret[0].([]review.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockReviewSourceMockRecorder) GetByUserID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockReviewSource)(nil).GetByUserID), arg0, arg1)
}
