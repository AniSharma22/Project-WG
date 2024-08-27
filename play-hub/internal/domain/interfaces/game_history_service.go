package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type GameHistoryService interface {
	GetTotalGameHistory() ([]entities.GameHistory, error)
	GetResultsToUpdate() ([]entities.GameHistory, error)
	UpdateResult(result string, slotId primitive.ObjectID) error
}
