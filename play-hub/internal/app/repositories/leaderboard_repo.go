package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	interfaces "project2/internal/domain/interfaces/repository"
	"project2/internal/models"
)

type leaderboardRepo struct {
	db *sql.DB
}

func NewLeaderboardRepo(db *sql.DB) interfaces.LeaderboardRepository {
	return &leaderboardRepo{db: db}
}

func (r *leaderboardRepo) FetchGameLeaderboard(ctx context.Context, gameID uuid.UUID) ([]models.Leaderboard, error) {
	query := `
		SELECT u.username, l.score
		FROM leaderboard l
		INNER JOIN users u ON l.user_id = u.user_id
		WHERE l.game_id = $1
		ORDER BY l.score DESC
	`
	rows, err := r.db.QueryContext(ctx, query, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch game leaderboard: %w", err)
	}
	defer rows.Close()

	var leaderboard []models.Leaderboard
	for rows.Next() {
		var entry models.Leaderboard
		if err := rows.Scan(&entry.UserName, &entry.Score); err != nil {
			return nil, fmt.Errorf("failed to scan leaderboard row: %w", err)
		}
		leaderboard = append(leaderboard, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over leaderboard: %w", err)
	}

	return leaderboard, nil
}

// FetchUserGameStats retrieves a user's stats for a specific game.
func (r *leaderboardRepo) FetchUserGameStats(ctx context.Context, userID, gameID uuid.UUID) (*entities.Leaderboard, error) {
	query := `SELECT score_id, user_id, game_id, wins, losses, score, created_at FROM leaderboard WHERE user_id = $1 AND game_id = $2`
	row := r.db.QueryRowContext(ctx, query, userID, gameID)

	var stats entities.Leaderboard
	err := row.Scan(&stats.ScoreID, &stats.UserID, &stats.GameID, &stats.Wins, &stats.Losses, &stats.Score, &stats.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No stats found for this user and game
		}
		return nil, fmt.Errorf("failed to fetch user game stats: %w", err)
	}

	return &stats, nil
}

// FetchUserOverallStats retrieves a user's overall stats across all games.
func (r *leaderboardRepo) FetchUserOverallStats(ctx context.Context, userID uuid.UUID) ([]entities.Leaderboard, error) {
	query := `SELECT score_id, user_id, game_id, wins, losses, score, created_at FROM leaderboard WHERE user_id = $1 ORDER BY score DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user overall stats: %w", err)
	}
	defer rows.Close()

	var stats []entities.Leaderboard
	for rows.Next() {
		var entry entities.Leaderboard
		if err := rows.Scan(&entry.ScoreID, &entry.UserID, &entry.GameID, &entry.Wins, &entry.Losses, &entry.Score, &entry.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan stats row: %w", err)
		}
		stats = append(stats, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over stats: %w", err)
	}

	return stats, nil
}

// UpdateUserGameStats updates a user's stats for a specific game.
func (r *leaderboardRepo) UpdateUserGameStats(ctx context.Context, leaderboard *entities.Leaderboard) error {
	query := `UPDATE leaderboard SET wins = $1, losses = $2, score = $3 WHERE user_id = $4 AND game_id = $5`
	_, err := r.db.ExecContext(ctx, query, leaderboard.Wins, leaderboard.Losses, leaderboard.Score, leaderboard.UserID, leaderboard.GameID)
	if err != nil {
		return fmt.Errorf("failed to update user game stats: %w", err)
	}
	return nil
}
