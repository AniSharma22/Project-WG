package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"time"
)

type SlotRepository interface {
	GetSlotsByDate(date string, gameId string) ([]entities.Slot, error)
	GetSlotByDateAndTime(date string, gameId string, time time.Time) (*entities.Slot, error)
	BookSlot(userId primitive.ObjectID, date string, gameId string, time time.Time) error
}
