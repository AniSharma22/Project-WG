package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"project2/internal/models"
	"time"
)

type BookingService struct {
	bookRepo    repository_interfaces.BookingRepository
	SlotService service_interfaces.SlotService
	GameService service_interfaces.GameService
}

func NewBookingService(bookRepo repository_interfaces.BookingRepository, slotService service_interfaces.SlotService, gameService service_interfaces.GameService) service_interfaces.BookingService {
	return &BookingService{
		bookRepo:    bookRepo,
		SlotService: slotService,
		GameService: gameService,
	}
}

// MakeBooking creates a new booking
func (b *BookingService) MakeBooking(ctx context.Context, userID uuid.UUID, slotID uuid.UUID) error {
	// Fetch the slot to check its current status
	slot, err := b.SlotService.GetSlotByID(ctx, slotID)
	if err != nil {
		return fmt.Errorf("failed to get slot details: %w", err)
	}

	if slot.IsBooked {
		return fmt.Errorf("slot is already booked")
	}
	if slot.StartTime.Before(time.Now()) {
		return fmt.Errorf("slot has already passed")
	}
	booking, err := b.bookRepo.FetchBookingBySlotAndUserId(ctx, slotID, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch booking: %w", err)
	}
	if booking.BookingId != uuid.Nil {
		return fmt.Errorf("user is already booked in this slot")
	}

	newBooking := &entities.Booking{
		SlotID: slotID,
		UserID: userID,
	}
	// Create a new booking
	_, err = b.bookRepo.CreateBooking(ctx, newBooking)
	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}

	game, err := b.GameService.GetGameByID(ctx, slot.GameID)
	if err != nil {
		return fmt.Errorf("failed to get game details: %w", err)
	}

	bookings, err := b.bookRepo.FetchBookingsBySlotID(ctx, slotID)
	if err != nil {
		return fmt.Errorf("failed to fetch bookings: %w", err)
	}
	if len(bookings) == game.MaxPlayers {
		// Update the slot status to booked
		err = b.SlotService.MarkSlotAsBooked(ctx, slotID)
		if err != nil {
			return fmt.Errorf("failed to update slot status: %w", err)
		}
	}

	return nil
}

// GetUpcomingBookings retrieves all upcoming bookings for a given user.
func (b *BookingService) GetUpcomingBookings(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	return b.bookRepo.FetchUpcomingBookingsByUserID(ctx, userID)
}

// GetBookingsToUpdateResult retrieves all bookings which have pending result update
func (b *BookingService) GetBookingsToUpdateResult(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	return b.bookRepo.FetchBookingsToUpdateResult(ctx, userID)
}

// UpdateBookingResult updates the result of a particular booking with either win or loss
func (b *BookingService) UpdateBookingResult(ctx context.Context, bookingId uuid.UUID, result string) error {
	return b.bookRepo.UpdateBookingResult(ctx, bookingId, result)
}

// GetSlotBookedUsers retrieves and returns the list of all the booked users of a slot
func (b *BookingService) GetSlotBookedUsers(ctx context.Context, slotId uuid.UUID) ([]string, error) {
	return b.bookRepo.FetchSlotBookedUsers(ctx, slotId)
}

// GetBookingByUserAndSlotID retrieves the slot based on the user and slot ID
func (b *BookingService) GetBookingByUserAndSlotID(ctx context.Context, userID uuid.UUID, slotID uuid.UUID) (models.Bookings, error) {
	return b.bookRepo.FetchBookingBySlotAndUserId(ctx, slotID, userID)
}
