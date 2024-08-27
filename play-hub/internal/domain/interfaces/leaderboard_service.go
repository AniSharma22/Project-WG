package interfaces

import "project2/internal/domain/entities"

type LeaderboardService interface {
	GetOverallLeaderboard() ([]entities.User, error)
}
