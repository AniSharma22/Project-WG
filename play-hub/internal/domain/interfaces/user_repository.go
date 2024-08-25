package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	AddLoss(userId primitive.ObjectID) error
	AddWin(userId primitive.ObjectID) error
	GetAllUsers() ([]entities.User, error)
	EmailAlreadyExists(email string) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserById(userId primitive.ObjectID) (*entities.User, error)
	GetPendingInvites(email string) ([]primitive.ObjectID, error)
	DeleteInvite(slotId primitive.ObjectID) error
	AddToInvites(userId primitive.ObjectID, slotId primitive.ObjectID) error
	GetAllUsersByScore() ([]entities.User, error)
}
