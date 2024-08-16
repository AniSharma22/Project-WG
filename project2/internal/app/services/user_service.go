package services

import (
	"errors"
	"fmt"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"project2/pkg/utils"
)

type UserService struct {
	userRepo interfaces.UserRepository // Injecting the UserRepository interface
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Signup(user *entities.User) error {
	// Set user role
	user.Role = "user"

	// Generate UUID for the user
	uuid, err := utils.GetUuid()
	if err != nil {
		fmt.Println("Error generating UUID: ", err)
		return err
	}
	user.UserId = uuid

	// Hash the user's password
	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		fmt.Println("Error generating password hash: ", err)
		return err
	}
	user.Password = hashedPassword

	// Save the user in the repository
	err = s.userRepo.CreateUser(user)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return err
	}

	// Store the entry in the global map
	globals.UsersMap[user.UserId] = *user

	return nil
}

func (s *UserService) Login(email string, password string) (*entities.User, error) {
	user, exists := globals.UsersMap[email]
	if !exists {
		return nil, errors.New("user not found")
	}

	isValidUser := utils.VerifyPassword(user.Password, password)
	if !isValidUser {
		return nil, errors.New("wrong password")
	}

	globals.ActiveUser = user.Email

	return &user, nil
}
