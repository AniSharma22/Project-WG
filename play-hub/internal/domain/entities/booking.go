package entities

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	BookingID uuid.UUID `json:"booking_id" db:"booking_id"`
	SlotID    uuid.UUID `json:"slot_id" db:"slot_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Result    string    `json:"result" db:"result"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
