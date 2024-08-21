package interfaces

import (
	"project2/internal/domain/entities"
	"time"
)

type SlotRepository interface {
	GetSlotsByDate(date string, gameId string) ([]entities.SlotStats, error)
	GetSlotByDateAndTime(date string, gameId string, time time.Time) (entities.SlotStats, error)
	BookSlot(user entities.User, date string, gameId string, time time.Time) error
	GetPendingInvites(user entities.User, date string) ([]entities.Invites, error)
}
