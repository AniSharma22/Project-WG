package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *entities.Booking) (uuid.UUID, error)
	FetchBookingByID(ctx context.Context, id uuid.UUID) (*entities.Booking, error)
	DeleteBookingByID(ctx context.Context, id uuid.UUID) error
	FetchBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Booking, error)
	FetchBookingsBySlotID(ctx context.Context, slotID uuid.UUID) ([]entities.Booking, error)
	FetchUpcomingBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error)
	UpdateBookingResult(ctx context.Context, bookingId uuid.UUID, result string) error
	FetchBookingsToUpdateResult(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error)
	FetchSlotBookedUsers(ctx context.Context, slotId uuid.UUID) ([]string, error)
	FetchBookingBySlotAndUserId(ctx context.Context, slotId uuid.UUID, userID uuid.UUID) (models.Bookings, error)
}
