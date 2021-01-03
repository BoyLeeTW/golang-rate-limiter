// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/interfaces/memcached.go

// Package mock is a generated GoMock package.
package mock

import (
	memcache "github.com/bradfitz/gomemcache/memcache"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMemcachedClientInterface is a mock of MemcachedClientInterface interface.
type MockMemcachedClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMemcachedClientInterfaceMockRecorder
}

// MockMemcachedClientInterfaceMockRecorder is the mock recorder for MockMemcachedClientInterface.
type MockMemcachedClientInterfaceMockRecorder struct {
	mock *MockMemcachedClientInterface
}

// NewMockMemcachedClientInterface creates a new mock instance.
func NewMockMemcachedClientInterface(ctrl *gomock.Controller) *MockMemcachedClientInterface {
	mock := &MockMemcachedClientInterface{ctrl: ctrl}
	mock.recorder = &MockMemcachedClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemcachedClientInterface) EXPECT() *MockMemcachedClientInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMemcachedClientInterface) Add(item *memcache.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMemcachedClientInterfaceMockRecorder) Add(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMemcachedClientInterface)(nil).Add), item)
}

// Increment mocks base method.
func (m *MockMemcachedClientInterface) Increment(key string, delta uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Increment", key, delta)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment.
func (mr *MockMemcachedClientInterfaceMockRecorder) Increment(key, delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockMemcachedClientInterface)(nil).Increment), key, delta)
}

// MockMemcachedServiceInterface is a mock of MemcachedServiceInterface interface.
type MockMemcachedServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMemcachedServiceInterfaceMockRecorder
}

// MockMemcachedServiceInterfaceMockRecorder is the mock recorder for MockMemcachedServiceInterface.
type MockMemcachedServiceInterfaceMockRecorder struct {
	mock *MockMemcachedServiceInterface
}

// NewMockMemcachedServiceInterface creates a new mock instance.
func NewMockMemcachedServiceInterface(ctrl *gomock.Controller) *MockMemcachedServiceInterface {
	mock := &MockMemcachedServiceInterface{ctrl: ctrl}
	mock.recorder = &MockMemcachedServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemcachedServiceInterface) EXPECT() *MockMemcachedServiceInterfaceMockRecorder {
	return m.recorder
}

// IncrementOrAdd mocks base method.
func (m *MockMemcachedServiceInterface) IncrementOrAdd(key string, delta uint64, expiration int32) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementOrAdd", key, delta, expiration)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrementOrAdd indicates an expected call of IncrementOrAdd.
func (mr *MockMemcachedServiceInterfaceMockRecorder) IncrementOrAdd(key, delta, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementOrAdd", reflect.TypeOf((*MockMemcachedServiceInterface)(nil).IncrementOrAdd), key, delta, expiration)
}