// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package stores is a generated GoMock package.
package stores

import (
	gomock "github.com/golang/mock/gomock"
	models "go_lang/Assignment/user-curd/models"
	reflect "reflect"
)

// MockCrud is a mock of Crud interface
type MockCrud struct {
	ctrl     *gomock.Controller
	recorder *MockCrudMockRecorder
}

// MockCrudMockRecorder is the mock recorder for MockCrud
type MockCrudMockRecorder struct {
	mock *MockCrud
}

// NewMockCrud creates a new mock instance
func NewMockCrud(ctrl *gomock.Controller) *MockCrud {
	mock := &MockCrud{ctrl: ctrl}
	mock.recorder = &MockCrudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrud) EXPECT() *MockCrudMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCrud) Create(name, email, phone string, age int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, email, phone, age)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCrudMockRecorder) Create(name, email, phone, age interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCrud)(nil).Create), name, email, phone, age)
}

// ReadOne mocks base method
func (m *MockCrud) ReadOne(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOne", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOne indicates an expected call of ReadOne
func (mr *MockCrudMockRecorder) ReadOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOne", reflect.TypeOf((*MockCrud)(nil).ReadOne), id)
}

// ReadAll mocks base method
func (m *MockCrud) ReadAll() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll
func (mr *MockCrudMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockCrud)(nil).ReadAll))
}

// Update mocks base method
func (m *MockCrud) Update(id int, name, email, phone string, age int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, name, email, phone, age)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCrudMockRecorder) Update(id, name, email, phone, age interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCrud)(nil).Update), id, name, email, phone, age)
}

// Delete mocks base method
func (m *MockCrud) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCrudMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCrud)(nil).Delete), id)
}
