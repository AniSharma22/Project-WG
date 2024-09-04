package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (uuid.UUID, error)
	FetchUserByEmail(ctx context.Context, email string) (*entities.User, error)
	FetchUserById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FetchAllUsers(ctx context.Context) ([]entities.User, error)
	EmailAlreadyExists(ctx context.Context, email string) bool
	FetchUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
