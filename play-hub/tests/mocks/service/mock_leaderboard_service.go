package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type MockLeaderboardService struct {
}

func NewMockLeaderboardService() *MockLeaderboardService {
	return &MockLeaderboardService{}
}

func (l *MockLeaderboardService) GetOverallLeaderboard() ([]entities.User, error) {
	return []entities.User{
		{
			ID:           primitive.NewObjectID(),
			OverallScore: 0,
			Email:        "test.test@watchguard.com",
		},
	}, nil
}
