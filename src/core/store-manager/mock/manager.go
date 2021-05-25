// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	manifest "github.com/ConsenSysQuorum/quorum-key-manager/src/core/manifest"
	entities "github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
	eth1 "github.com/ConsenSysQuorum/quorum-key-manager/src/store/eth1"
	keys "github.com/ConsenSysQuorum/quorum-key-manager/src/store/keys"
	secrets "github.com/ConsenSysQuorum/quorum-key-manager/src/store/secrets"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStoreManager is a mock of StoreManager interface
type MockStoreManager struct {
	ctrl     *gomock.Controller
	recorder *MockStoreManagerMockRecorder
}

// MockStoreManagerMockRecorder is the mock recorder for MockStoreManager
type MockStoreManagerMockRecorder struct {
	mock *MockStoreManager
}

// NewMockStoreManager creates a new mock instance
func NewMockStoreManager(ctrl *gomock.Controller) *MockStoreManager {
	mock := &MockStoreManager{ctrl: ctrl}
	mock.recorder = &MockStoreManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreManager) EXPECT() *MockStoreManagerMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockStoreManager) Load(ctx context.Context, mnfsts ...*manifest.Manifest) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range mnfsts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Load", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockStoreManagerMockRecorder) Load(ctx interface{}, mnfsts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, mnfsts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStoreManager)(nil).Load), varargs...)
}

// GetSecretStore mocks base method
func (m *MockStoreManager) GetSecretStore(ctx context.Context, name string) (secrets.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretStore", ctx, name)
	ret0, _ := ret[0].(secrets.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretStore indicates an expected call of GetSecretStore
func (mr *MockStoreManagerMockRecorder) GetSecretStore(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretStore", reflect.TypeOf((*MockStoreManager)(nil).GetSecretStore), ctx, name)
}

// GetKeyStore mocks base method
func (m *MockStoreManager) GetKeyStore(ctx context.Context, name string) (keys.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyStore", ctx, name)
	ret0, _ := ret[0].(keys.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyStore indicates an expected call of GetKeyStore
func (mr *MockStoreManagerMockRecorder) GetKeyStore(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyStore", reflect.TypeOf((*MockStoreManager)(nil).GetKeyStore), ctx, name)
}

// GetEth1Store mocks base method
func (m *MockStoreManager) GetEth1Store(ctx context.Context, name string) (eth1.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEth1Store", ctx, name)
	ret0, _ := ret[0].(eth1.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEth1Store indicates an expected call of GetEth1Store
func (mr *MockStoreManagerMockRecorder) GetEth1Store(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEth1Store", reflect.TypeOf((*MockStoreManager)(nil).GetEth1Store), ctx, name)
}

// GetEth1StoreByAddr mocks base method
func (m *MockStoreManager) GetEth1StoreByAddr(ctx context.Context, addr common.Address) (eth1.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEth1StoreByAddr", ctx, addr)
	ret0, _ := ret[0].(eth1.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEth1StoreByAddr indicates an expected call of GetEth1StoreByAddr
func (mr *MockStoreManagerMockRecorder) GetEth1StoreByAddr(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEth1StoreByAddr", reflect.TypeOf((*MockStoreManager)(nil).GetEth1StoreByAddr), ctx, addr)
}

// List mocks base method
func (m *MockStoreManager) List(ctx context.Context, kind manifest.Kind) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, kind)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreManagerMockRecorder) List(ctx, kind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStoreManager)(nil).List), ctx, kind)
}

// ListAllAccounts mocks base method
func (m *MockStoreManager) ListAllAccounts(arg0 context.Context) ([]*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllAccounts indicates an expected call of ListAllAccounts
func (mr *MockStoreManagerMockRecorder) ListAllAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStoreManager)(nil).ListAllAccounts), arg0)
}
