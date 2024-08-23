package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"project2/internal/domain/entities"
)

type SlotRepository interface {
	GetSlotsByDate(date string, gameId primitive.ObjectID) ([]entities.Slot, error)
	GetSlotByDateAndTime(date string, gameId primitive.ObjectID, startTime string) (*entities.Slot, error)
	BookSlot(userId primitive.ObjectID, date string, gameId primitive.ObjectID, startTime string) error
	InsertSlot(slot entities.Slot) (*mongo.InsertOneResult, error)
	GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error)
}
