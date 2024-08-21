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
}
