// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

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

// DeleteById mocks base method.
func (m *MockStore) DeleteById(id int, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockStoreMockRecorder) DeleteById(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockStore)(nil).DeleteById), id, ctx)
}

// GetAllProducts mocks base method.
func (m *MockStore) GetAllProducts(ctx *gofr.Context) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", ctx)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockStoreMockRecorder) GetAllProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockStore)(nil).GetAllProducts), ctx)
}

// GetById mocks base method.
func (m *MockStore) GetById(id int, ctx *gofr.Context) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id, ctx)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockStoreMockRecorder) GetById(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockStore)(nil).GetById), id, ctx)
}

// InsertProduct mocks base method.
func (m *MockStore) InsertProduct(product models.Product, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProduct", product, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProduct indicates an expected call of InsertProduct.
func (mr *MockStoreMockRecorder) InsertProduct(product, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProduct", reflect.TypeOf((*MockStore)(nil).InsertProduct), product, ctx)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(product models.Product, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", product, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(product, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), product, ctx)
}
