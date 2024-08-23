package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type GameRepository interface {
	GetGameByID(gameId primitive.ObjectID) (*entities.Game, error)
	GetAllGames() ([]entities.Game, error)
	CreateGame(game *entities.Game) error
	DeleteGame(gameId primitive.ObjectID) error
}
