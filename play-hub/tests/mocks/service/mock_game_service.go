package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

// MockGameService is a mock implementation of the GameService interface
type MockGameService struct{}

// NewMockGameService creates a new instance of MockGameService
func NewMockGameService() *MockGameService {
	return &MockGameService{}
}

// GetGameByID returns a mock game with the given ID
func (s *MockGameService) GetGameByID(gameID primitive.ObjectID) (*entities.Game, error) {
	return &entities.Game{
		ID:          gameID,
		Name:        "test",
		MaxCapacity: 1,
	}, nil
}

// GetAllGames returns a slice containing a single mock game
func (s *MockGameService) GetAllGames() ([]entities.Game, error) {
	return []entities.Game{
		{
			ID:          primitive.NewObjectID(),
			Name:        "test",
			MaxCapacity: 1,
		},
	}, nil
}

// CreateGame mocks the creation of a game without actually performing any action
func (s *MockGameService) CreateGame(name string, maxPlayers int) error {
	return nil
}

func (s *MockGameService) DeleteGame(gameId primitive.ObjectID) error {
	return nil
}
