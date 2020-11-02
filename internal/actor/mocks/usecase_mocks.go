// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_actor is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockActorUseCase is a mock of ActorUseCase interface
type MockActorUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockActorUseCaseMockRecorder
}

// MockActorUseCaseMockRecorder is the mock recorder for MockActorUseCase
type MockActorUseCaseMockRecorder struct {
	mock *MockActorUseCase
}

// NewMockActorUseCase creates a new mock instance
func NewMockActorUseCase(ctrl *gomock.Controller) *MockActorUseCase {
	mock := &MockActorUseCase{ctrl: ctrl}
	mock.recorder = &MockActorUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActorUseCase) EXPECT() *MockActorUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockActorUseCase) Create(actor *models.Actor) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", actor)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockActorUseCaseMockRecorder) Create(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActorUseCase)(nil).Create), actor)
}

// Get mocks base method
func (m *MockActorUseCase) Get(id uint64) (*models.Actor, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*models.Actor)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockActorUseCaseMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockActorUseCase)(nil).Get), id)
}

// Change mocks base method
func (m *MockActorUseCase) Change(newActor *models.Actor) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Change", newActor)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Change indicates an expected call of Change
func (mr *MockActorUseCaseMockRecorder) Change(newActor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change", reflect.TypeOf((*MockActorUseCase)(nil).Change), newActor)
}

// DeleteById mocks base method
func (m *MockActorUseCase) DeleteById(id uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById
func (mr *MockActorUseCaseMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockActorUseCase)(nil).DeleteById), id)
}