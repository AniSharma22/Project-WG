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

func NewGameService(gameRepo interfaces.GameRepository) *GameService {
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

//func (s *GameService) CreateGame(game *entities.Game) error {
//	game.GameId, _ = utils.GetUuid()
//	return s.gameRepo.CreateGame(game)
//}

//func (s *GameService) DeleteGame(gameId string) error {
//	return s.gameRepo.DeleteGame(gameId)
//}
