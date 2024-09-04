// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Dell\projects2\watchguard-proj\Project-WG\play-hub\internal\domain\interfaces\repository\booking_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "project2/internal/domain/entities"
	models "project2/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockBookingRepository is a mock of BookingRepository interface.
type MockBookingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookingRepositoryMockRecorder
}

// MockBookingRepositoryMockRecorder is the mock recorder for MockBookingRepository.
type MockBookingRepositoryMockRecorder struct {
	mock *MockBookingRepository
}

// NewMockBookingRepository creates a new mock instance.
func NewMockBookingRepository(ctrl *gomock.Controller) *MockBookingRepository {
	mock := &MockBookingRepository{ctrl: ctrl}
	mock.recorder = &MockBookingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingRepository) EXPECT() *MockBookingRepositoryMockRecorder {
	return m.recorder
}

// CreateBooking mocks base method.
func (m *MockBookingRepository) CreateBooking(ctx context.Context, booking *entities.Booking) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBooking", ctx, booking)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBooking indicates an expected call of CreateBooking.
func (mr *MockBookingRepositoryMockRecorder) CreateBooking(ctx, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBooking", reflect.TypeOf((*MockBookingRepository)(nil).CreateBooking), ctx, booking)
}

// DeleteBookingByID mocks base method.
func (m *MockBookingRepository) DeleteBookingByID(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookingByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookingByID indicates an expected call of DeleteBookingByID.
func (mr *MockBookingRepositoryMockRecorder) DeleteBookingByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookingByID", reflect.TypeOf((*MockBookingRepository)(nil).DeleteBookingByID), ctx, id)
}

// FetchBookingByID mocks base method.
func (m *MockBookingRepository) FetchBookingByID(ctx context.Context, id uuid.UUID) (*entities.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookingByID", ctx, id)
	ret0, _ := ret[0].(*entities.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookingByID indicates an expected call of FetchBookingByID.
func (mr *MockBookingRepositoryMockRecorder) FetchBookingByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookingByID", reflect.TypeOf((*MockBookingRepository)(nil).FetchBookingByID), ctx, id)
}

// FetchBookingBySlotAndUserId mocks base method.
func (m *MockBookingRepository) FetchBookingBySlotAndUserId(ctx context.Context, slotId, userID uuid.UUID) (models.Bookings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookingBySlotAndUserId", ctx, slotId, userID)
	ret0, _ := ret[0].(models.Bookings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookingBySlotAndUserId indicates an expected call of FetchBookingBySlotAndUserId.
func (mr *MockBookingRepositoryMockRecorder) FetchBookingBySlotAndUserId(ctx, slotId, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookingBySlotAndUserId", reflect.TypeOf((*MockBookingRepository)(nil).FetchBookingBySlotAndUserId), ctx, slotId, userID)
}

// FetchBookingsBySlotID mocks base method.
func (m *MockBookingRepository) FetchBookingsBySlotID(ctx context.Context, slotID uuid.UUID) ([]entities.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookingsBySlotID", ctx, slotID)
	ret0, _ := ret[0].([]entities.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookingsBySlotID indicates an expected call of FetchBookingsBySlotID.
func (mr *MockBookingRepositoryMockRecorder) FetchBookingsBySlotID(ctx, slotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookingsBySlotID", reflect.TypeOf((*MockBookingRepository)(nil).FetchBookingsBySlotID), ctx, slotID)
}

// FetchBookingsByUserID mocks base method.
func (m *MockBookingRepository) FetchBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookingsByUserID", ctx, userID)
	ret0, _ := ret[0].([]entities.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookingsByUserID indicates an expected call of FetchBookingsByUserID.
func (mr *MockBookingRepositoryMockRecorder) FetchBookingsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookingsByUserID", reflect.TypeOf((*MockBookingRepository)(nil).FetchBookingsByUserID), ctx, userID)
}

// FetchBookingsToUpdateResult mocks base method.
func (m *MockBookingRepository) FetchBookingsToUpdateResult(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookingsToUpdateResult", ctx, userID)
	ret0, _ := ret[0].([]models.Bookings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookingsToUpdateResult indicates an expected call of FetchBookingsToUpdateResult.
func (mr *MockBookingRepositoryMockRecorder) FetchBookingsToUpdateResult(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookingsToUpdateResult", reflect.TypeOf((*MockBookingRepository)(nil).FetchBookingsToUpdateResult), ctx, userID)
}

// FetchSlotBookedUsers mocks base method.
func (m *MockBookingRepository) FetchSlotBookedUsers(ctx context.Context, slotId uuid.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSlotBookedUsers", ctx, slotId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSlotBookedUsers indicates an expected call of FetchSlotBookedUsers.
func (mr *MockBookingRepositoryMockRecorder) FetchSlotBookedUsers(ctx, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSlotBookedUsers", reflect.TypeOf((*MockBookingRepository)(nil).FetchSlotBookedUsers), ctx, slotId)
}

// FetchUpcomingBookingsByUserID mocks base method.
func (m *MockBookingRepository) FetchUpcomingBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUpcomingBookingsByUserID", ctx, userID)
	ret0, _ := ret[0].([]models.Bookings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUpcomingBookingsByUserID indicates an expected call of FetchUpcomingBookingsByUserID.
func (mr *MockBookingRepositoryMockRecorder) FetchUpcomingBookingsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUpcomingBookingsByUserID", reflect.TypeOf((*MockBookingRepository)(nil).FetchUpcomingBookingsByUserID), ctx, userID)
}

// UpdateBookingResult mocks base method.
func (m *MockBookingRepository) UpdateBookingResult(ctx context.Context, bookingId uuid.UUID, result string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookingResult", ctx, bookingId, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBookingResult indicates an expected call of UpdateBookingResult.
func (mr *MockBookingRepositoryMockRecorder) UpdateBookingResult(ctx, bookingId, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookingResult", reflect.TypeOf((*MockBookingRepository)(nil).UpdateBookingResult), ctx, bookingId, result)
}
