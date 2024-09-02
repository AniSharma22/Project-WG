package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	"time"
)

type SlotRepository interface {
	FetchSlotByID(ctx context.Context, id uuid.UUID) (*entities.Slot, error)
	CreateSlot(ctx context.Context, slot *entities.Slot) (uuid.UUID, error)
	DeleteSlotByID(ctx context.Context, id uuid.UUID) error
	FetchSlotsByDate(ctx context.Context, date time.Time) ([]entities.Slot, error)
	FetchSlotByDateAndTime(ctx context.Context, date time.Time, startTime time.Time) (*entities.Slot, error)
	FetchSlotsByGameID(ctx context.Context, gameID uuid.UUID) ([]entities.Slot, error)
	FetchSlotsByGameIDAndDate(ctx context.Context, gameID uuid.UUID, date time.Time) ([]entities.Slot, error)
	UpdateSlotStatus(ctx context.Context, slotID uuid.UUID, isBooked bool) error
	//GetSlotsByDate(date time.Time, gameId primitive.ObjectID) ([]entities.Slot, error)
	//GetSlotByDateAndTime(date time.Time, gameId primitive.ObjectID, startTime time.Time) (*entities.Slot, error)
	//BookSlot(userId primitive.ObjectID, slotId primitive.ObjectID) error
	//InsertSlot(slot entities.Slot) (*mongo.InsertOneResult, error)
	//GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error)
	//GetUpcomingBookedSlots(userId primitive.ObjectID) ([]entities.Slot, error)
	//AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error
}