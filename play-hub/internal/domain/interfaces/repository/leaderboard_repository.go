package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type LeaderboardRepository interface {
	// FetchGameLeaderboard retrieves the leaderboard for a specific game.
	FetchGameLeaderboard(ctx context.Context, gameID uuid.UUID) ([]models.Leaderboard, error)

	// FetchUserGameStats retrieves a user's stats for a specific game.
	FetchUserGameStats(ctx context.Context, userID, gameID uuid.UUID) (*entities.Leaderboard, error)

	// FetchUserOverallStats retrieves a user's overall stats across all games.
	FetchUserOverallStats(ctx context.Context, userID uuid.UUID) ([]entities.Leaderboard, error)

	// UpdateUserGameStats updates a user's stats for a specific game.
	UpdateUserGameStats(ctx context.Context, leaderboard *entities.Leaderboard) error
}
