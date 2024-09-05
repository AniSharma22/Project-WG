package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
)

type GameRepository interface {
	FetchGameByID(ctx context.Context, id uuid.UUID) (*entities.Game, error)
	FetchAllGames(ctx context.Context) ([]entities.Game, error)
	CreateGame(ctx context.Context, game *entities.Game) (uuid.UUID, error)
	DeleteGame(ctx context.Context, id uuid.UUID) error
	UpdateGameStatus(ctx context.Context, gameID uuid.UUID, status bool) error
}
