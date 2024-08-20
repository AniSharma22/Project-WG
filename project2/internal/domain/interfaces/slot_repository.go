package interfaces

import (
	"project2/internal/domain/entities"
	"time"
)

type SlotRepository interface {
	GetSlotsByDate(date time.Time) ([]entities.Slot, error)
	BookSlot(slotID string, user entities.User) error
	GetPendingInvites(userID string) ([]entities.Slot, error)
	UpdateSlot(slot entities.Slot) error
}
