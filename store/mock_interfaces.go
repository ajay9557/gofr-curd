// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package store is a generated GoMock package.
package store

import (
	models "gofr-curd/models"
	reflect "reflect"

	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStore) Create(ctx *gofr.Context, pd models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, pd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(ctx, pd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), ctx, pd)
}

// Delete mocks base method.
func (m *MockStore) Delete(ctx *gofr.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockStore) Get(ctx *gofr.Context) ([]*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx)
}

// GetByID mocks base method.
func (m *MockStore) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStoreMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStore)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *MockStore) Update(ctx *gofr.Context, pd models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, pd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStoreMockRecorder) Update(ctx, pd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), ctx, pd)
}
