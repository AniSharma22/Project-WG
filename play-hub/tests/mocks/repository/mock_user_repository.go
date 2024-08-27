// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\anisharma\GolandProjects\goprac\Project-WG\play-hub\internal\domain\interfaces\user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "project2/internal/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddLoss mocks base method.
func (m *MockUserRepository) AddLoss(userId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLoss", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLoss indicates an expected call of AddLoss.
func (mr *MockUserRepositoryMockRecorder) AddLoss(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLoss", reflect.TypeOf((*MockUserRepository)(nil).AddLoss), userId)
}

// AddToInvites mocks base method.
func (m *MockUserRepository) AddToInvites(userId, slotId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToInvites", userId, slotId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToInvites indicates an expected call of AddToInvites.
func (mr *MockUserRepositoryMockRecorder) AddToInvites(userId, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToInvites", reflect.TypeOf((*MockUserRepository)(nil).AddToInvites), userId, slotId)
}

// AddWin mocks base method.
func (m *MockUserRepository) AddWin(userId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWin", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWin indicates an expected call of AddWin.
func (mr *MockUserRepositoryMockRecorder) AddWin(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWin", reflect.TypeOf((*MockUserRepository)(nil).AddWin), userId)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// DeleteInvite mocks base method.
func (m *MockUserRepository) DeleteInvite(slotId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvite", slotId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInvite indicates an expected call of DeleteInvite.
func (mr *MockUserRepositoryMockRecorder) DeleteInvite(slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvite", reflect.TypeOf((*MockUserRepository)(nil).DeleteInvite), slotId)
}

// EmailAlreadyExists mocks base method.
func (m *MockUserRepository) EmailAlreadyExists(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmailAlreadyExists", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmailAlreadyExists indicates an expected call of EmailAlreadyExists.
func (mr *MockUserRepositoryMockRecorder) EmailAlreadyExists(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmailAlreadyExists", reflect.TypeOf((*MockUserRepository)(nil).EmailAlreadyExists), email)
}

// GetAllUsers mocks base method.
func (m *MockUserRepository) GetAllUsers() ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepositoryMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepository)(nil).GetAllUsers))
}

// GetAllUsersByScore mocks base method.
func (m *MockUserRepository) GetAllUsersByScore() ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersByScore")
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersByScore indicates an expected call of GetAllUsersByScore.
func (mr *MockUserRepositoryMockRecorder) GetAllUsersByScore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersByScore", reflect.TypeOf((*MockUserRepository)(nil).GetAllUsersByScore))
}

// GetPendingInvites mocks base method.
func (m *MockUserRepository) GetPendingInvites(email string) ([]primitive.ObjectID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingInvites", email)
	ret0, _ := ret[0].([]primitive.ObjectID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingInvites indicates an expected call of GetPendingInvites.
func (mr *MockUserRepositoryMockRecorder) GetPendingInvites(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingInvites", reflect.TypeOf((*MockUserRepository)(nil).GetPendingInvites), email)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// GetUserById mocks base method.
func (m *MockUserRepository) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", userId)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepositoryMockRecorder) GetUserById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepository)(nil).GetUserById), userId)
}
