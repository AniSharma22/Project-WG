package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"time"
)

type MockSlotService struct {
}

func NewMockSlotService() *MockSlotService {
	return &MockSlotService{}
}

// Mock implementation for GetGameSlots
func (s *MockSlotService) GetGameSlots(game *entities.Game) ([]entities.Slot, error) {
	// Return mock data
	return []entities.Slot{
		{
			ID:          primitive.NewObjectID(),
			GameID:      game.ID,
			StartTime:   time.Now().Add(1 * time.Hour),
			BookedUsers: []primitive.ObjectID{},
		},
	}, nil
}

// Mock implementation for BookSlot
func (s *MockSlotService) BookSlot(game *entities.Game, slot *entities.Slot) error {
	return nil
}

// Mock implementation for InviteToSlot
func (s *MockSlotService) InviteToSlot(userId primitive.ObjectID, game *entities.Game, slot *entities.Slot) error {
	return nil
}

// Mock implementation for GetSlotById
func (s *MockSlotService) GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error) {
	return &entities.Slot{
		ID:        slotId,
		GameID:    primitive.NewObjectID(),
		StartTime: time.Now().Add(1 * time.Hour),
	}, nil
}

// Mock implementation for GetUpcomingBookedSlots
func (s *MockSlotService) GetUpcomingBookedSlots() ([]entities.Slot, error) {
	return []entities.Slot{
		{
			ID:          primitive.NewObjectID(),
			GameID:      primitive.NewObjectID(),
			StartTime:   time.Now(),
			EndTime:     time.Now().Add(20 * time.Minute),
			BookedUsers: []primitive.ObjectID{},
			Results:     []entities.Result{},
		},
	}, nil
}

// Mock implementation for AddResultToSlot
func (s *MockSlotService) AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error {
	return nil
}
