package entities

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	NotificationID uuid.UUID `json:"notification_id" db:"notification_id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	Message        string    `json:"message" db:"message"`
	IsRead         bool      `json:"is_read" db:"is_read"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
