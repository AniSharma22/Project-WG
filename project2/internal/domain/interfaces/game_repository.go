package interfaces

import "project2/internal/domain/entities"

type GameRepository interface {
	GetGameByID(gameID string) (*entities.Game, error)
	GetAllGames() ([]entities.Game, error)
	CreateGame(game *entities.Game) error
	DeleteGame(gameID string) error
}
