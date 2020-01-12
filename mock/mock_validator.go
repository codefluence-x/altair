// Code generated by MockGen. DO NOT EDIT.
// Source: core/validator.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/codefluence-x/altair/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockOauthValidator is a mock of OauthValidator interface
type MockOauthValidator struct {
	ctrl     *gomock.Controller
	recorder *MockOauthValidatorMockRecorder
}

// MockOauthValidatorMockRecorder is the mock recorder for MockOauthValidator
type MockOauthValidatorMockRecorder struct {
	mock *MockOauthValidator
}

// NewMockOauthValidator creates a new mock instance
func NewMockOauthValidator(ctrl *gomock.Controller) *MockOauthValidator {
	mock := &MockOauthValidator{ctrl: ctrl}
	mock.recorder = &MockOauthValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthValidator) EXPECT() *MockOauthValidatorMockRecorder {
	return m.recorder
}

// ValidateApplication mocks base method
func (m *MockOauthValidator) ValidateApplication(ctx context.Context, data entity.OauthApplicationJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateApplication", ctx, data)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateApplication indicates an expected call of ValidateApplication
func (mr *MockOauthValidatorMockRecorder) ValidateApplication(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateApplication", reflect.TypeOf((*MockOauthValidator)(nil).ValidateApplication), ctx, data)
}
