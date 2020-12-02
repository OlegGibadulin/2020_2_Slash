// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_content is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContentUsecase is a mocks of ContentUsecase interface
type MockContentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockContentUsecaseMockRecorder
}

// MockContentUsecaseMockRecorder is the mocks recorder for MockContentUsecase
type MockContentUsecaseMockRecorder struct {
	mock *MockContentUsecase
}

// NewMockContentUsecase creates a new mocks instance
func NewMockContentUsecase(ctrl *gomock.Controller) *MockContentUsecase {
	mock := &MockContentUsecase{ctrl: ctrl}
	mock.recorder = &MockContentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContentUsecase) EXPECT() *MockContentUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockContentUsecase) Create(content *models.Content) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", content)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockContentUsecaseMockRecorder) Create(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContentUsecase)(nil).Create), content)
}

// Update mocks base method
func (m *MockContentUsecase) Update(newContentData *models.Content) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", newContentData)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockContentUsecaseMockRecorder) Update(newContentData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContentUsecase)(nil).Update), newContentData)
}

// UpdatePosters mocks base method
func (m *MockContentUsecase) UpdatePosters(content *models.Content, newPostersDir string) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePosters", content, newPostersDir)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// UpdatePosters indicates an expected call of UpdatePosters
func (mr *MockContentUsecaseMockRecorder) UpdatePosters(content, newPostersDir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePosters", reflect.TypeOf((*MockContentUsecase)(nil).UpdatePosters), content, newPostersDir)
}

// DeleteByID mocks base method
func (m *MockContentUsecase) DeleteByID(contentID uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", contentID)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockContentUsecaseMockRecorder) DeleteByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockContentUsecase)(nil).DeleteByID), contentID)
}

// GetByID mocks base method
func (m *MockContentUsecase) GetByID(contentID uint64) (*models.Content, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", contentID)
	ret0, _ := ret[0].(*models.Content)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockContentUsecaseMockRecorder) GetByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockContentUsecase)(nil).GetByID), contentID)
}

// GetFullByID mocks base method
func (m *MockContentUsecase) GetFullByID(contentID uint64) (*models.Content, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullByID", contentID)
	ret0, _ := ret[0].(*models.Content)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetFullByID indicates an expected call of GetFullByID
func (mr *MockContentUsecaseMockRecorder) GetFullByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullByID", reflect.TypeOf((*MockContentUsecase)(nil).GetFullByID), contentID)
}

// FillContent mocks base method
func (m *MockContentUsecase) FillContent(content *models.Content) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillContent", content)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// FillContent indicates an expected call of FillContent
func (mr *MockContentUsecaseMockRecorder) FillContent(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillContent", reflect.TypeOf((*MockContentUsecase)(nil).FillContent), content)
}

// GetCountriesByID mocks base method
func (m *MockContentUsecase) GetCountriesByID(contentID uint64) ([]*models.Country, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountriesByID", contentID)
	ret0, _ := ret[0].([]*models.Country)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetCountriesByID indicates an expected call of GetCountriesByID
func (mr *MockContentUsecaseMockRecorder) GetCountriesByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountriesByID", reflect.TypeOf((*MockContentUsecase)(nil).GetCountriesByID), contentID)
}

// GetGenresByID mocks base method
func (m *MockContentUsecase) GetGenresByID(contentID uint64) ([]*models.Genre, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenresByID", contentID)
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetGenresByID indicates an expected call of GetGenresByID
func (mr *MockContentUsecaseMockRecorder) GetGenresByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenresByID", reflect.TypeOf((*MockContentUsecase)(nil).GetGenresByID), contentID)
}

// GetActorsByID mocks base method
func (m *MockContentUsecase) GetActorsByID(contentID uint64) ([]*models.Actor, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByID", contentID)
	ret0, _ := ret[0].([]*models.Actor)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetActorsByID indicates an expected call of GetActorsByID
func (mr *MockContentUsecaseMockRecorder) GetActorsByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByID", reflect.TypeOf((*MockContentUsecase)(nil).GetActorsByID), contentID)
}

// GetDirectorsByID mocks base method
func (m *MockContentUsecase) GetDirectorsByID(contentID uint64) ([]*models.Director, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectorsByID", contentID)
	ret0, _ := ret[0].([]*models.Director)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetDirectorsByID indicates an expected call of GetDirectorsByID
func (mr *MockContentUsecaseMockRecorder) GetDirectorsByID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectorsByID", reflect.TypeOf((*MockContentUsecase)(nil).GetDirectorsByID), contentID)
}
