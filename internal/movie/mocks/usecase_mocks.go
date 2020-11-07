// Code generated by MockGen. DO NOT EDIT.
// Source: internal/movie/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMovieUsecase is a mock of MovieUsecase interface
type MockMovieUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMovieUsecaseMockRecorder
}

// MockMovieUsecaseMockRecorder is the mock recorder for MockMovieUsecase
type MockMovieUsecaseMockRecorder struct {
	mock *MockMovieUsecase
}

// NewMockMovieUsecase creates a new mock instance
func NewMockMovieUsecase(ctrl *gomock.Controller) *MockMovieUsecase {
	mock := &MockMovieUsecase{ctrl: ctrl}
	mock.recorder = &MockMovieUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieUsecase) EXPECT() *MockMovieUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMovieUsecase) Create(movie *models.Movie) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", movie)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMovieUsecaseMockRecorder) Create(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMovieUsecase)(nil).Create), movie)
}

// UpdateVideo mocks base method
func (m *MockMovieUsecase) UpdateVideo(movie *models.Movie, newVideoPath string) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVideo", movie, newVideoPath)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// UpdateVideo indicates an expected call of UpdateVideo
func (mr *MockMovieUsecaseMockRecorder) UpdateVideo(movie, newVideoPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVideo", reflect.TypeOf((*MockMovieUsecase)(nil).UpdateVideo), movie, newVideoPath)
}

// DeleteByID mocks base method
func (m *MockMovieUsecase) DeleteByID(movieID uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", movieID)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockMovieUsecaseMockRecorder) DeleteByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockMovieUsecase)(nil).DeleteByID), movieID)
}

// GetByID mocks base method
func (m *MockMovieUsecase) GetByID(movieID uint64) (*models.Movie, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", movieID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockMovieUsecaseMockRecorder) GetByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMovieUsecase)(nil).GetByID), movieID)
}

// GetWithContentByID mocks base method
func (m *MockMovieUsecase) GetWithContentByID(movieID uint64) (*models.Movie, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithContentByID", movieID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetWithContentByID indicates an expected call of GetWithContentByID
func (mr *MockMovieUsecaseMockRecorder) GetWithContentByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithContentByID", reflect.TypeOf((*MockMovieUsecase)(nil).GetWithContentByID), movieID)
}

// GetFullByID mocks base method
func (m *MockMovieUsecase) GetFullByID(movieID uint64) (*models.Movie, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullByID", movieID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetFullByID indicates an expected call of GetFullByID
func (mr *MockMovieUsecaseMockRecorder) GetFullByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullByID", reflect.TypeOf((*MockMovieUsecase)(nil).GetFullByID), movieID)
}

// GetByContentID mocks base method
func (m *MockMovieUsecase) GetByContentID(contentID uint64) (*models.Movie, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByContentID", contentID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByContentID indicates an expected call of GetByContentID
func (mr *MockMovieUsecaseMockRecorder) GetByContentID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByContentID", reflect.TypeOf((*MockMovieUsecase)(nil).GetByContentID), contentID)
}

// ListByGenre mocks base method
func (m *MockMovieUsecase) ListByGenre(genreID uint64) ([]*models.Movie, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByGenre", genreID)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListByGenre indicates an expected call of ListByGenre
func (mr *MockMovieUsecaseMockRecorder) ListByGenre(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByGenre", reflect.TypeOf((*MockMovieUsecase)(nil).ListByGenre), genreID)
}