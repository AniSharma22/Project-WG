package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type UserService interface {
	Signup(user *entities.User) error
	EmailAlreadyExists(email string) bool
	Login(email string, password []byte) (*entities.User, error)
	GetUserById(userId primitive.ObjectID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetPendingInvites() ([]models.Invite, error)
	AcceptInvite(slotId primitive.ObjectID) error
	RejectInvite(slotId primitive.ObjectID) error
	AddResult(userId primitive.ObjectID, result string) error
	GetAllUsersByScore() ([]entities.User, error)
}
