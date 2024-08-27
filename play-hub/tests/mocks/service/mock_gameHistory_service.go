package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type MockGameHistoryService struct {
}

func NewMockGameHistoryService() *MockGameHistoryService {
	return &MockGameHistoryService{}
}

func (gh *MockGameHistoryService) GetTotalGameHistory() ([]entities.GameHistory, error) {
	gameHistories := []entities.GameHistory{
		{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			GameID: primitive.NewObjectID(),
			SlotID: primitive.NewObjectID(),
			Result: "win",
		},
	}

	return gameHistories, nil
}
func (gh *MockGameHistoryService) GetResultsToUpdate() ([]entities.GameHistory, error) {
	gameHistories := []entities.GameHistory{
		{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			GameID: primitive.NewObjectID(),
			SlotID: primitive.NewObjectID(),
			Result: "win",
		},
	}

	return gameHistories, nil
}
func (gh *MockGameHistoryService) UpdateResult(result string, slotId primitive.ObjectID) error {
	return nil
}
