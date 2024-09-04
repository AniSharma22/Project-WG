package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type LeaderboardRepository interface {
	FetchGameLeaderboard(ctx context.Context, gameID uuid.UUID) ([]models.Leaderboard, error)
	FetchUserGameStats(ctx context.Context, userID, gameID uuid.UUID) (*entities.Leaderboard, error)
	FetchUserOverallStats(ctx context.Context, userID uuid.UUID) ([]entities.Leaderboard, error)
	UpdateUserGameStats(ctx context.Context, leaderboard *entities.Leaderboard) error
}
