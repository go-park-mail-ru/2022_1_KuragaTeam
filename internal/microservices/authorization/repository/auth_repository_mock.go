// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/microservices/authorization/repository.go

// Package repository is a generated GoMock package.
package repository

import (
	proto "myapp/internal/microservices/authorization/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockStorage) CreateUser(data *proto.SignUpData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorageMockRecorder) CreateUser(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser), data)
}

// DeleteSession mocks base method.
func (m *MockStorage) DeleteSession(session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStorageMockRecorder) DeleteSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStorage)(nil).DeleteSession), session)
}

// GetUserId mocks base method.
func (m *MockStorage) GetUserId(session string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId", session)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId.
func (mr *MockStorageMockRecorder) GetUserId(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockStorage)(nil).GetUserId), session)
}

// IsUserExists mocks base method.
func (m *MockStorage) IsUserExists(data *proto.LogInData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExists", data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserExists indicates an expected call of IsUserExists.
func (mr *MockStorageMockRecorder) IsUserExists(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExists", reflect.TypeOf((*MockStorage)(nil).IsUserExists), data)
}

// IsUserUnique mocks base method.
func (m *MockStorage) IsUserUnique(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserUnique", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserUnique indicates an expected call of IsUserUnique.
func (mr *MockStorageMockRecorder) IsUserUnique(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserUnique", reflect.TypeOf((*MockStorage)(nil).IsUserUnique), email)
}

// StoreSession mocks base method.
func (m *MockStorage) StoreSession(userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSession", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreSession indicates an expected call of StoreSession.
func (mr *MockStorageMockRecorder) StoreSession(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSession", reflect.TypeOf((*MockStorage)(nil).StoreSession), userID)
}