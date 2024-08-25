package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"project2/internal/domain/entities"
	"time"
)

type SlotRepository interface {
	GetSlotsByDate(date time.Time, gameId primitive.ObjectID) ([]entities.Slot, error)
	GetSlotByDateAndTime(date time.Time, gameId primitive.ObjectID, startTime time.Time) (*entities.Slot, error)
	BookSlot(userId primitive.ObjectID, slotId primitive.ObjectID) error
	InsertSlot(slot entities.Slot) (*mongo.InsertOneResult, error)
	GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error)
	GetUpcomingBookedSlots(userId primitive.ObjectID) ([]entities.Slot, error)
	AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error
}
