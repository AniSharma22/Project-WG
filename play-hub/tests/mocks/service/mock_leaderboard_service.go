// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\leaderboard_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLeaderboardService is a mock of LeaderboardService interface.
type MockLeaderboardService struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderboardServiceMockRecorder
}

// MockLeaderboardServiceMockRecorder is the mock recorder for MockLeaderboardService.
type MockLeaderboardServiceMockRecorder struct {
	mock *MockLeaderboardService
}

// NewMockLeaderboardService creates a new mock instance.
func NewMockLeaderboardService(ctrl *gomock.Controller) *MockLeaderboardService {
	mock := &MockLeaderboardService{ctrl: ctrl}
	mock.recorder = &MockLeaderboardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderboardService) EXPECT() *MockLeaderboardServiceMockRecorder {
	return m.recorder
}

// GetOverallLeaderboard mocks base method.
func (m *MockLeaderboardService) GetOverallLeaderboard() ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOverallLeaderboard")
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOverallLeaderboard indicates an expected call of GetOverallLeaderboard.
func (mr *MockLeaderboardServiceMockRecorder) GetOverallLeaderboard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOverallLeaderboard", reflect.TypeOf((*MockLeaderboardService)(nil).GetOverallLeaderboard))
}
