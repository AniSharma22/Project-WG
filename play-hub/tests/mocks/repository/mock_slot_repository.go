// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\slot_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockSlotRepository is a mock of SlotRepository interface.
type MockSlotRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSlotRepositoryMockRecorder
}

// MockSlotRepositoryMockRecorder is the mock recorder for MockSlotRepository.
type MockSlotRepositoryMockRecorder struct {
	mock *MockSlotRepository
}

// NewMockSlotRepository creates a new mock instance.
func NewMockSlotRepository(ctrl *gomock.Controller) *MockSlotRepository {
	mock := &MockSlotRepository{ctrl: ctrl}
	mock.recorder = &MockSlotRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlotRepository) EXPECT() *MockSlotRepositoryMockRecorder {
	return m.recorder
}

// AddResultToSlot mocks base method.
func (m *MockSlotRepository) AddResultToSlot(userId, slotId primitive.ObjectID, result string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResultToSlot", userId, slotId, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResultToSlot indicates an expected call of AddResultToSlot.
func (mr *MockSlotRepositoryMockRecorder) AddResultToSlot(userId, slotId, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResultToSlot", reflect.TypeOf((*MockSlotRepository)(nil).AddResultToSlot), userId, slotId, result)
}

// BookSlot mocks base method.
func (m *MockSlotRepository) BookSlot(userId, slotId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookSlot", userId, slotId)
	ret0, _ := ret[0].(error)
	return ret0
}

// BookSlot indicates an expected call of BookSlot.
func (mr *MockSlotRepositoryMockRecorder) BookSlot(userId, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookSlot", reflect.TypeOf((*MockSlotRepository)(nil).BookSlot), userId, slotId)
}

// GetSlotByDateAndTime mocks base method.
func (m *MockSlotRepository) GetSlotByDateAndTime(date time.Time, gameId primitive.ObjectID, startTime time.Time) (*entities.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotByDateAndTime", date, gameId, startTime)
	ret0, _ := ret[0].(*entities.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotByDateAndTime indicates an expected call of GetSlotByDateAndTime.
func (mr *MockSlotRepositoryMockRecorder) GetSlotByDateAndTime(date, gameId, startTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotByDateAndTime", reflect.TypeOf((*MockSlotRepository)(nil).GetSlotByDateAndTime), date, gameId, startTime)
}

// GetSlotById mocks base method.
func (m *MockSlotRepository) GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotById", slotId)
	ret0, _ := ret[0].(*entities.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotById indicates an expected call of GetSlotById.
func (mr *MockSlotRepositoryMockRecorder) GetSlotById(slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotById", reflect.TypeOf((*MockSlotRepository)(nil).GetSlotById), slotId)
}

// GetSlotsByDate mocks base method.
func (m *MockSlotRepository) GetSlotsByDate(date time.Time, gameId primitive.ObjectID) ([]entities.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotsByDate", date, gameId)
	ret0, _ := ret[0].([]entities.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotsByDate indicates an expected call of GetSlotsByDate.
func (mr *MockSlotRepositoryMockRecorder) GetSlotsByDate(date, gameId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotsByDate", reflect.TypeOf((*MockSlotRepository)(nil).GetSlotsByDate), date, gameId)
}

// GetUpcomingBookedSlots mocks base method.
func (m *MockSlotRepository) GetUpcomingBookedSlots(userId primitive.ObjectID) ([]entities.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingBookedSlots", userId)
	ret0, _ := ret[0].([]entities.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingBookedSlots indicates an expected call of GetUpcomingBookedSlots.
func (mr *MockSlotRepositoryMockRecorder) GetUpcomingBookedSlots(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingBookedSlots", reflect.TypeOf((*MockSlotRepository)(nil).GetUpcomingBookedSlots), userId)
}

// InsertSlot mocks base method.
func (m *MockSlotRepository) InsertSlot(slot entities.Slot) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSlot", slot)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertSlot indicates an expected call of InsertSlot.
func (mr *MockSlotRepositoryMockRecorder) InsertSlot(slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSlot", reflect.TypeOf((*MockSlotRepository)(nil).InsertSlot), slot)
}
