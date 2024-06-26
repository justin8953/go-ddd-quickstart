// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen --source=db.go --destination=./mock/mock_db.go
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	db "go-ddd-quickstart/internal/pkg/db"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDbRepo is a mock of DbRepo interface.
type MockDbRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDbRepoMockRecorder
}

// MockDbRepoMockRecorder is the mock recorder for MockDbRepo.
type MockDbRepoMockRecorder struct {
	mock *MockDbRepo
}

// NewMockDbRepo creates a new mock instance.
func NewMockDbRepo(ctrl *gomock.Controller) *MockDbRepo {
	mock := &MockDbRepo{ctrl: ctrl}
	mock.recorder = &MockDbRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbRepo) EXPECT() *MockDbRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDbRepo) Create(item db.IItem) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDbRepoMockRecorder) Create(item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDbRepo)(nil).Create), item)
}

// Delete mocks base method.
func (m *MockDbRepo) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDbRepoMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDbRepo)(nil).Delete), id)
}

// List mocks base method.
func (m *MockDbRepo) List(filter map[string]any) ([]db.IItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filter)
	ret0, _ := ret[0].([]db.IItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDbRepoMockRecorder) List(filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDbRepo)(nil).List), filter)
}

// Retrieve mocks base method.
func (m *MockDbRepo) Retrieve(id string) (db.IItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", id)
	ret0, _ := ret[0].(db.IItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockDbRepoMockRecorder) Retrieve(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockDbRepo)(nil).Retrieve), id)
}

// Update mocks base method.
func (m *MockDbRepo) Update(id string, item db.IItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDbRepoMockRecorder) Update(id, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDbRepo)(nil).Update), id, item)
}
