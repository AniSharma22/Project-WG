package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type GameService interface {
	GetGameByID(gameID primitive.ObjectID) (*entities.Game, error)
	GetAllGames() ([]entities.Game, error)
	CreateGame(name string, maxPlayers int) error
	DeleteGame(gameId primitive.ObjectID) error
}
