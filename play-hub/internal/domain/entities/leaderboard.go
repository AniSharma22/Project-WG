package entities

import (
	"github.com/google/uuid"
	"time"
)

type Leaderboard struct {
	ScoreID   uuid.UUID `json:"score_id" db:"score_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	GameID    uuid.UUID `json:"game_id" db:"game_id"`
	Wins      int       `json:"wins" db:"wins"`
	Losses    int       `json:"losses" db:"losses"`
	Score     float64   `json:"score" db:"score"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
