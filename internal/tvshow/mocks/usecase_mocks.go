// Code generated by MockGen. DO NOT EDIT.
// Source: internal/tvshow/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTVShowUsecase is a mock of TVShowUsecase interface
type MockTVShowUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTVShowUsecaseMockRecorder
}

// MockTVShowUsecaseMockRecorder is the mock recorder for MockTVShowUsecase
type MockTVShowUsecaseMockRecorder struct {
	mock *MockTVShowUsecase
}

// NewMockTVShowUsecase creates a new mock instance
func NewMockTVShowUsecase(ctrl *gomock.Controller) *MockTVShowUsecase {
	mock := &MockTVShowUsecase{ctrl: ctrl}
	mock.recorder = &MockTVShowUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTVShowUsecase) EXPECT() *MockTVShowUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTVShowUsecase) Create(tvshow *models.TVShow) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tvshow)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTVShowUsecaseMockRecorder) Create(tvshow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTVShowUsecase)(nil).Create), tvshow)
}

// GetByID mocks base method
func (m *MockTVShowUsecase) GetByID(tvshowID uint64) (*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", tvshowID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTVShowUsecaseMockRecorder) GetByID(tvshowID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTVShowUsecase)(nil).GetByID), tvshowID)
}

// GetFullByID mocks base method
func (m *MockTVShowUsecase) GetFullByID(tvshowID, curUserID uint64) (*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullByID", tvshowID, curUserID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetFullByID indicates an expected call of GetFullByID
func (mr *MockTVShowUsecaseMockRecorder) GetFullByID(tvshowID, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullByID", reflect.TypeOf((*MockTVShowUsecase)(nil).GetFullByID), tvshowID, curUserID)
}

// GetByContentID mocks base method
func (m *MockTVShowUsecase) GetByContentID(contentID uint64) (*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByContentID", contentID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByContentID indicates an expected call of GetByContentID
func (mr *MockTVShowUsecaseMockRecorder) GetByContentID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByContentID", reflect.TypeOf((*MockTVShowUsecase)(nil).GetByContentID), contentID)
}

// ListByParams mocks base method
func (m *MockTVShowUsecase) ListByParams(params *models.ContentFilter, pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByParams", params, pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListByParams indicates an expected call of ListByParams
func (mr *MockTVShowUsecaseMockRecorder) ListByParams(params, pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByParams", reflect.TypeOf((*MockTVShowUsecase)(nil).ListByParams), params, pgnt, curUserID)
}

// ListLatest mocks base method
func (m *MockTVShowUsecase) ListLatest(pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLatest", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListLatest indicates an expected call of ListLatest
func (mr *MockTVShowUsecaseMockRecorder) ListLatest(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLatest", reflect.TypeOf((*MockTVShowUsecase)(nil).ListLatest), pgnt, curUserID)
}

// ListByRating mocks base method
func (m *MockTVShowUsecase) ListByRating(pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByRating", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListByRating indicates an expected call of ListByRating
func (mr *MockTVShowUsecaseMockRecorder) ListByRating(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByRating", reflect.TypeOf((*MockTVShowUsecase)(nil).ListByRating), pgnt, curUserID)
}
