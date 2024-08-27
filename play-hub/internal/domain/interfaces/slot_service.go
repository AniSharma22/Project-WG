package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type SlotService interface {
	GetGameSlots(game *entities.Game) ([]entities.Slot, error)
	BookSlot(game *entities.Game, slot *entities.Slot) error
	InviteToSlot(userId primitive.ObjectID, game *entities.Game, slot *entities.Slot) error
	GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error)
	GetUpcomingBookedSlots() ([]entities.Slot, error)
	AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error
}
