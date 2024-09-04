package entities

import (
	"github.com/google/uuid"
	"time"
)

type Invitation struct {
	InvitationID   uuid.UUID `json:"invitation_id" db:"invitation_id"`
	InvitingUserID uuid.UUID `json:"inviting_user_id" db:"inviting_user_id"`
	InvitedUserID  uuid.UUID `json:"invited_user_id" db:"invited_user_id"`
	SlotID         uuid.UUID `json:"slot_id" db:"slot_id"`
	Status         string    `json:"status" db:"status"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
