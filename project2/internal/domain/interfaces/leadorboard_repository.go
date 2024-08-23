package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type LeaderboardRepository interface {
	GetGameLeaderboard(gameId primitive.ObjectID) ([]entities.Leaderboard, error)
	GetOverallLeaderboard() ([]entities.Leaderboard, error)
	AddOrUpdateLeaderboardEntry(entry *entities.Leaderboard) error
}
