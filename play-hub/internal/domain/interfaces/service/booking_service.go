package service_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/models"
)

type BookingService interface {
	MakeBooking(ctx context.Context, userID uuid.UUID, slotID uuid.UUID) error
	GetUpcomingBookings(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error)
	GetBookingsToUpdateResult(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error)
	UpdateBookingResult(ctx context.Context, bookingId uuid.UUID, result string) error
	GetSlotBookedUsers(ctx context.Context, slotId uuid.UUID) ([]string, error)
	GetBookingByUserAndSlotID(ctx context.Context, userID uuid.UUID, slotID uuid.UUID) (models.Bookings, error)
}
