// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\game_history_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockGameHistoryRepository is a mock of GameHistoryRepository interface.
type MockGameHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameHistoryRepositoryMockRecorder
}

// MockGameHistoryRepositoryMockRecorder is the mock recorder for MockGameHistoryRepository.
type MockGameHistoryRepositoryMockRecorder struct {
	mock *MockGameHistoryRepository
}

// NewMockGameHistoryRepository creates a new mock instance.
func NewMockGameHistoryRepository(ctrl *gomock.Controller) *MockGameHistoryRepository {
	mock := &MockGameHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockGameHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameHistoryRepository) EXPECT() *MockGameHistoryRepositoryMockRecorder {
	return m.recorder
}

// AddGameHistory mocks base method.
func (m *MockGameHistoryRepository) AddGameHistory(history *entities.GameHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGameHistory", history)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGameHistory indicates an expected call of AddGameHistory.
func (mr *MockGameHistoryRepositoryMockRecorder) AddGameHistory(history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGameHistory", reflect.TypeOf((*MockGameHistoryRepository)(nil).AddGameHistory), history)
}

// FindGameHistoryByID mocks base method.
func (m *MockGameHistoryRepository) FindGameHistoryByID(historyID primitive.ObjectID) (*entities.GameHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGameHistoryByID", historyID)
	ret0, _ := ret[0].(*entities.GameHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGameHistoryByID indicates an expected call of FindGameHistoryByID.
func (mr *MockGameHistoryRepositoryMockRecorder) FindGameHistoryByID(historyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGameHistoryByID", reflect.TypeOf((*MockGameHistoryRepository)(nil).FindGameHistoryByID), historyID)
}

// GetAllGameHistories mocks base method.
func (m *MockGameHistoryRepository) GetAllGameHistories() ([]entities.GameHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGameHistories")
	ret0, _ := ret[0].([]entities.GameHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGameHistories indicates an expected call of GetAllGameHistories.
func (mr *MockGameHistoryRepositoryMockRecorder) GetAllGameHistories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGameHistories", reflect.TypeOf((*MockGameHistoryRepository)(nil).GetAllGameHistories))
}

// GetCurrentDayHistory mocks base method.
func (m *MockGameHistoryRepository) GetCurrentDayHistory(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentDayHistory", userId)
	ret0, _ := ret[0].([]entities.GameHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentDayHistory indicates an expected call of GetCurrentDayHistory.
func (mr *MockGameHistoryRepositoryMockRecorder) GetCurrentDayHistory(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentDayHistory", reflect.TypeOf((*MockGameHistoryRepository)(nil).GetCurrentDayHistory), userId)
}

// GetResultsToUpdate mocks base method.
func (m *MockGameHistoryRepository) GetResultsToUpdate(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsToUpdate", userId)
	ret0, _ := ret[0].([]entities.GameHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResultsToUpdate indicates an expected call of GetResultsToUpdate.
func (mr *MockGameHistoryRepositoryMockRecorder) GetResultsToUpdate(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsToUpdate", reflect.TypeOf((*MockGameHistoryRepository)(nil).GetResultsToUpdate), userId)
}

// GetUserGameHistory mocks base method.
func (m *MockGameHistoryRepository) GetUserGameHistory(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGameHistory", userId)
	ret0, _ := ret[0].([]entities.GameHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGameHistory indicates an expected call of GetUserGameHistory.
func (mr *MockGameHistoryRepositoryMockRecorder) GetUserGameHistory(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGameHistory", reflect.TypeOf((*MockGameHistoryRepository)(nil).GetUserGameHistory), userId)
}

// RemoveGameHistory mocks base method.
func (m *MockGameHistoryRepository) RemoveGameHistory(historyID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGameHistory", historyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGameHistory indicates an expected call of RemoveGameHistory.
func (mr *MockGameHistoryRepositoryMockRecorder) RemoveGameHistory(historyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGameHistory", reflect.TypeOf((*MockGameHistoryRepository)(nil).RemoveGameHistory), historyID)
}

// UpdateResult mocks base method.
func (m *MockGameHistoryRepository) UpdateResult(result string, slotId, userID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResult", result, slotId, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResult indicates an expected call of UpdateResult.
func (mr *MockGameHistoryRepositoryMockRecorder) UpdateResult(result, slotId, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResult", reflect.TypeOf((*MockGameHistoryRepository)(nil).UpdateResult), result, slotId, userID)
}
