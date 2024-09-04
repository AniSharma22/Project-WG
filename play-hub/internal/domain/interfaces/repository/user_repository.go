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
	//AddLoss(userId primitive.ObjectID) error
	//AddWin(userId primitive.ObjectID) error
	//GetAllUsers() ([]entities.User, error)
	//GetUserByEmail(email string) (*entities.User, error)
	//GetUserById(userId primitive.ObjectID) (*entities.User, error)
	//GetPendingInvites(email string) ([]primitive.ObjectID, error)
	//DeleteInvite(slotId primitive.ObjectID) error
	//AddToInvites(userId primitive.ObjectID, slotId primitive.ObjectID) error
	//GetAllUsersByScore() ([]entities.User, error)
}
