// Code generated by MockGen. DO NOT EDIT.
// Source: internal/genre/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGenreUsecase is a mock of GenreUsecase interface
type MockGenreUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockGenreUsecaseMockRecorder
}

// MockGenreUsecaseMockRecorder is the mock recorder for MockGenreUsecase
type MockGenreUsecaseMockRecorder struct {
	mock *MockGenreUsecase
}

// NewMockGenreUsecase creates a new mock instance
func NewMockGenreUsecase(ctrl *gomock.Controller) *MockGenreUsecase {
	mock := &MockGenreUsecase{ctrl: ctrl}
	mock.recorder = &MockGenreUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenreUsecase) EXPECT() *MockGenreUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGenreUsecase) Create(genre *models.Genre) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", genre)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGenreUsecaseMockRecorder) Create(genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGenreUsecase)(nil).Create), genre)
}

// UpdateByID mocks base method
func (m *MockGenreUsecase) UpdateByID(genreID uint64, newGenreData *models.Genre) (*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", genreID, newGenreData)
	ret0, _ := ret[0].(*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockGenreUsecaseMockRecorder) UpdateByID(genreID, newGenreData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockGenreUsecase)(nil).UpdateByID), genreID, newGenreData)
}

// DeleteByID mocks base method
func (m *MockGenreUsecase) DeleteByID(genreID uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", genreID)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockGenreUsecaseMockRecorder) DeleteByID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockGenreUsecase)(nil).DeleteByID), genreID)
}

// GetByID mocks base method
func (m *MockGenreUsecase) GetByID(genreID uint64) (*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", genreID)
	ret0, _ := ret[0].(*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockGenreUsecaseMockRecorder) GetByID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGenreUsecase)(nil).GetByID), genreID)
}

// GetByName mocks base method
func (m *MockGenreUsecase) GetByName(name string) (*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockGenreUsecaseMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockGenreUsecase)(nil).GetByName), name)
}

// List mocks base method
func (m *MockGenreUsecase) List() ([]*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGenreUsecaseMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGenreUsecase)(nil).List))
}

// ListByID mocks base method
func (m *MockGenreUsecase) ListByID(genresID []uint64) ([]*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByID", genresID)
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListByID indicates an expected call of ListByID
func (mr *MockGenreUsecaseMockRecorder) ListByID(genresID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByID", reflect.TypeOf((*MockGenreUsecase)(nil).ListByID), genresID)
}
