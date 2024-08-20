package interfaces

import "project2/internal/domain/entities"

type Leaderboard interface {
	GetGameLeaderboard(gameId string, users []entities.User) ([]entities.User, error)
	GetOverallLeaderboard(users []entities.User) ([]entities.User, error)
}
