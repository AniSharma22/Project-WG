package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type InvitationRepository interface {
	CreateInvitation(ctx context.Context, invitation *entities.Invitation) (uuid.UUID, error)
	DeleteInvitationByID(ctx context.Context, id uuid.UUID) error
	UpdateInvitationStatus(ctx context.Context, id uuid.UUID, status string) error
	FetchInvitationByID(ctx context.Context, id uuid.UUID) (*entities.Invitation, error)
	FetchUserInvitations(ctx context.Context, userID uuid.UUID) ([]entities.Invitation, error)
	FetchUserPendingInvitations(ctx context.Context, userID uuid.UUID) ([]models.Invitations, error)
	FetchInvitationByUserAndSlot(ctx context.Context, invitingUserID uuid.UUID, invitedUserID uuid.UUID, slotID uuid.UUID) (*entities.Invitation, error)
}
