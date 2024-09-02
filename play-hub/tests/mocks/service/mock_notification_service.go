// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\notification_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockNotificationService is a mock of NotificationService interface.
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService.
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance.
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// GetUserNotifications mocks base method.
func (m *MockNotificationService) GetUserNotifications(userId primitive.ObjectID) ([]entities.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserNotifications", userId)
	ret0, _ := ret[0].([]entities.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserNotifications indicates an expected call of GetUserNotifications.
func (mr *MockNotificationServiceMockRecorder) GetUserNotifications(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserNotifications", reflect.TypeOf((*MockNotificationService)(nil).GetUserNotifications), userId)
}