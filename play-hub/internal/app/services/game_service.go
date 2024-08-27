package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"sync"
)

type GameService struct {
	gameRepo interfaces.GameRepository
	gameWG   *sync.WaitGroup
}

func NewGameService(gameRepo interfaces.GameRepository) interfaces.GameService {
	return &GameService{
		gameRepo: gameRepo,
		gameWG:   &sync.WaitGroup{},
	}
}

func (s *GameService) GetGameByID(gameID primitive.ObjectID) (*entities.Game, error) {
	return s.gameRepo.GetGameByID(gameID)
}

func (s *GameService) GetAllGames() ([]entities.Game, error) {
	return s.gameRepo.GetAllGames()
}

func (s *GameService) CreateGame(name string, maxPlayers int) error {
	game := entities.Game{
		ID:          primitive.NewObjectID(),
		Name:        name,
		MaxCapacity: maxPlayers,
	}
	return s.gameRepo.CreateGame(&game)
}

func (s *GameService) DeleteGame(gameId primitive.ObjectID) error {
	return s.gameRepo.DeleteGame(gameId)
}
