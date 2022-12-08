// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/strangelove-ventures/packet-forward-middleware/v3/router/types (interfaces: BankKeeper)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(arg0 types.Context, arg1 string, arg2 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), arg0, arg1, arg2)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(arg0 types.Context, arg1, arg2 types.AccAddress, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), arg0, arg1, arg2, arg3)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(arg0 types.Context, arg1 types.AccAddress, arg2 string, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), arg0, arg1, arg2, arg3)
}