package entities

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	GameID     uuid.UUID `json:"game_id" db:"game_id"`
	GameName   string    `json:"game_name" db:"game_name"`
	MinPlayers int       `json:"min_players" db:"min_players"`
	MaxPlayers int       `json:"max_players" db:"max_players"`
	Instances  int       `json:"instances" db:"instances"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
