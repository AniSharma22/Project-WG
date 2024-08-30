// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\game_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockGameService is a mock of GameService interface.
type MockGameService struct {
	ctrl     *gomock.Controller
	recorder *MockGameServiceMockRecorder
}

// MockGameServiceMockRecorder is the mock recorder for MockGameService.
type MockGameServiceMockRecorder struct {
	mock *MockGameService
}

// NewMockGameService creates a new mock instance.
func NewMockGameService(ctrl *gomock.Controller) *MockGameService {
	mock := &MockGameService{ctrl: ctrl}
	mock.recorder = &MockGameServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameService) EXPECT() *MockGameServiceMockRecorder {
	return m.recorder
}

// CreateGame mocks base method.
func (m *MockGameService) CreateGame(name string, maxPlayers int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", name, maxPlayers)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockGameServiceMockRecorder) CreateGame(name, maxPlayers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockGameService)(nil).CreateGame), name, maxPlayers)
}

// DeleteGame mocks base method.
func (m *MockGameService) DeleteGame(gameId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGame", gameId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGame indicates an expected call of DeleteGame.
func (mr *MockGameServiceMockRecorder) DeleteGame(gameId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGame", reflect.TypeOf((*MockGameService)(nil).DeleteGame), gameId)
}

// GetAllGames mocks base method.
func (m *MockGameService) GetAllGames() ([]entities.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGames")
	ret0, _ := ret[0].([]entities.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGames indicates an expected call of GetAllGames.
func (mr *MockGameServiceMockRecorder) GetAllGames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGames", reflect.TypeOf((*MockGameService)(nil).GetAllGames))
}

// GetGameByID mocks base method.
func (m *MockGameService) GetGameByID(gameID primitive.ObjectID) (*entities.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByID", gameID)
	ret0, _ := ret[0].(*entities.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByID indicates an expected call of GetGameByID.
func (mr *MockGameServiceMockRecorder) GetGameByID(gameID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByID", reflect.TypeOf((*MockGameService)(nil).GetGameByID), gameID)
}
