package services

import (
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/utils"
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

func (s *GameService) GetGameByID(gameID string) (*entities.Game, error) {
	return s.gameRepo.GetGameByID(gameID)
}

func (s *GameService) GetAllGames() ([]entities.Game, error) {
	return s.gameRepo.GetAllGames()
}

func (s *GameService) CreateGame(game *entities.Game) error {
	game.GameId, _ = utils.GetUuid()
	return s.gameRepo.CreateGame(game)
}

func (s *GameService) DeleteGame(gameId string) error {
	return s.gameRepo.DeleteGame(gameId)
}
