// Code generated by MockGen. DO NOT EDIT.
// Source: block.go

// Package actionstest is a generated GoMock package.
package actionstest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBlock is a mock of Block interface
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockBlock) Run(ctx context.Context, protocol, network string, blk interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, protocol, network, blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockBlockMockRecorder) Run(ctx, protocol, network, blk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockBlock)(nil).Run), ctx, protocol, network, blk)
}
