package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"project2/internal/models"
	"project2/pkg/globals"
	"sync"
)

type InvitationService struct {
	invitationRepo repository_interfaces.InvitationRepository
	bookingService service_interfaces.BookingService
	slotService    service_interfaces.SlotService
	invitationWG   *sync.WaitGroup
}

func NewInvitationService(invitationRepo repository_interfaces.InvitationRepository, bookingService service_interfaces.BookingService, slotService service_interfaces.SlotService) service_interfaces.InvitationService {
	return &InvitationService{
		invitationRepo: invitationRepo,
		bookingService: bookingService,
		slotService:    slotService,
		invitationWG:   &sync.WaitGroup{},
	}
}

// MakeInvitation creates a new invitation.
func (s *InvitationService) MakeInvitation(ctx context.Context, invitingUserID, invitedUserID uuid.UUID, slotId uuid.UUID) (uuid.UUID, error) {

	// Check if the inviting user has already invited the same user for the same slot
	existingInvitation, err := s.invitationRepo.FetchInvitationByUserAndSlot(ctx, invitingUserID, invitedUserID, slotId)
	if err != nil {
		return uuid.Nil, err
	}
	if existingInvitation != nil {
		return uuid.Nil, errors.New("invitation already exists for this slot")
	}

	// Check if the slot is already booked
	slot, err := s.slotService.GetSlotByID(ctx, slotId)
	if err != nil {
		return uuid.Nil, err
	}
	if slot.IsBooked {
		return uuid.Nil, errors.New("slot is already booked")
	}

	invitation := &entities.Invitation{
		InvitingUserID: invitingUserID,
		InvitedUserID:  invitedUserID,
		SlotID:         slotId,
	}

	// Create the invitation in the repository
	invitationID, err := s.invitationRepo.CreateInvitation(ctx, invitation)
	if err != nil {
		return uuid.Nil, err
	}

	return invitationID, nil
}

// AcceptInvitation sets the status of an invitation to 'accepted'.
func (s *InvitationService) AcceptInvitation(ctx context.Context, invitationID uuid.UUID) error {
	invitation, err := s.invitationRepo.FetchInvitationByID(ctx, invitationID)
	if err != nil {
		return errors.New("failed to fetch invitation")
	}

	slot, err := s.slotService.GetSlotByID(ctx, invitation.SlotID)
	if err != nil {
		return errors.New("failed to fetch slot")
	}
	if slot.IsBooked {
		err = s.invitationRepo.DeleteInvitationByID(ctx, invitationID)
		if err != nil {
			return errors.New("failed to delete invitation")
		}
		return errors.New("slot is already booked")
	}

	booking, err := s.bookingService.GetBookingByUserAndSlotID(ctx, invitation.InvitedUserID, invitation.SlotID)
	if err != nil {
		return errors.New("failed to fetch invitation")
	}
	if booking.BookingId != uuid.Nil {
		err = s.invitationRepo.DeleteInvitationByID(ctx, invitationID)
		if err != nil {
			return errors.New("failed to delete invitation")
		}
		return errors.New("you already have this slot booked")
	}

	err = s.bookingService.MakeBooking(ctx, globals.ActiveUser, invitation.SlotID)
	if err != nil {
		return errors.New("failed to booking invitation")
	}
	err = s.invitationRepo.DeleteInvitationByID(ctx, invitationID)
	if err != nil {
		return errors.New("failed to delete invitation")
	}
	return nil
}

// RejectInvitation sets the status of an invitation to 'declined'.
func (s *InvitationService) RejectInvitation(ctx context.Context, invitationID uuid.UUID) error {
	err := s.invitationRepo.DeleteInvitationByID(ctx, invitationID)
	if err != nil {
		return errors.New("failed to reject invitation")
	}
	return nil
}

// GetAllPendingInvitations retrieves all pending invitations for a user.
func (s *InvitationService) GetAllPendingInvitations(ctx context.Context, userID uuid.UUID) ([]models.Invitations, error) {
	return s.invitationRepo.FetchUserPendingInvitations(ctx, userID)
}
