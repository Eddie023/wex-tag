// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eddie023/wex-tag/pkg/api (interfaces: ExchangeRateService)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_exchange_rate.go -package=mocks . ExchangeRateService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ent "github.com/eddie023/wex-tag/ent"
	service "github.com/eddie023/wex-tag/pkg/api/service"
	types "github.com/eddie023/wex-tag/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockExchangeRateService is a mock of ExchangeRateService interface.
type MockExchangeRateService struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRateServiceMockRecorder
}

// MockExchangeRateServiceMockRecorder is the mock recorder for MockExchangeRateService.
type MockExchangeRateServiceMockRecorder struct {
	mock *MockExchangeRateService
}

// NewMockExchangeRateService creates a new mock instance.
func NewMockExchangeRateService(ctrl *gomock.Controller) *MockExchangeRateService {
	mock := &MockExchangeRateService{ctrl: ctrl}
	mock.recorder = &MockExchangeRateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRateService) EXPECT() *MockExchangeRateServiceMockRecorder {
	return m.recorder
}

// ConvertCurrency mocks base method.
func (m *MockExchangeRateService) ConvertCurrency(arg0 service.ExchangeRatePayload, arg1 *ent.Transaction, arg2 service.ExchangeRateResponse) (types.GetPurchaseTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertCurrency", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.GetPurchaseTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertCurrency indicates an expected call of ConvertCurrency.
func (mr *MockExchangeRateServiceMockRecorder) ConvertCurrency(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertCurrency", reflect.TypeOf((*MockExchangeRateService)(nil).ConvertCurrency), arg0, arg1, arg2)
}

// GetExchangeRate mocks base method.
func (m *MockExchangeRateService) GetExchangeRate(arg0 context.Context, arg1 service.ExchangeRatePayload) (service.ExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(service.ExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockExchangeRateServiceMockRecorder) GetExchangeRate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockExchangeRateService)(nil).GetExchangeRate), arg0, arg1)
}
