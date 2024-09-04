package service_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/models"
)

type InvitationService interface {
	MakeInvitation(ctx context.Context, invitingUserID, invitedUserID uuid.UUID, slotId uuid.UUID) (uuid.UUID, error)
	AcceptInvitation(ctx context.Context, invitationID uuid.UUID) error
	RejectInvitation(ctx context.Context, invitationID uuid.UUID) error
	GetAllPendingInvitations(ctx context.Context, userID uuid.UUID) ([]models.Invitations, error)
}
