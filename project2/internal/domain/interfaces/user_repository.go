package interfaces

import "project2/internal/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	// Other methods for user repository
}
