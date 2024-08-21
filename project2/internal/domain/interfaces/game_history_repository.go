package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type GameHistoryRepository interface {
	AddGameHistory(history *entities.GameHistory) error
	RemoveGameHistory(historyID primitive.ObjectID) error
	FindGameHistoryByID(historyID primitive.ObjectID) (*entities.GameHistory, error)
	GetAllGameHistories() ([]entities.GameHistory, error)
}
