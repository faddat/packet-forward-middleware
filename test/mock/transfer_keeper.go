// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/strangelove-ventures/packet-forward-middleware/v2/router/types (interfaces: TransferKeeper)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	gomock "github.com/golang/mock/gomock"
)

// MockTransferKeeper is a mock of TransferKeeper interface.
type MockTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTransferKeeperMockRecorder
}

// MockTransferKeeperMockRecorder is the mock recorder for MockTransferKeeper.
type MockTransferKeeperMockRecorder struct {
	mock *MockTransferKeeper
}

// NewMockTransferKeeper creates a new mock instance.
func NewMockTransferKeeper(ctrl *gomock.Controller) *MockTransferKeeper {
	mock := &MockTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferKeeper) EXPECT() *MockTransferKeeperMockRecorder {
	return m.recorder
}

// SendTransfer mocks base method.
func (m *MockTransferKeeper) SendTransfer(arg0 types.Context, arg1, arg2 string, arg3 types.Coin, arg4 types.AccAddress, arg5 string, arg6 types0.Height, arg7 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransfer", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransfer indicates an expected call of SendTransfer.
func (mr *MockTransferKeeperMockRecorder) SendTransfer(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransfer", reflect.TypeOf((*MockTransferKeeper)(nil).SendTransfer), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
