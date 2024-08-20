package services

import (
	"errors"
	"fmt"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"sync"
)

type UserService struct {
	userRepo interfaces.UserRepository // dependency injection ??
	userWG   *sync.WaitGroup
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		userWG:   &sync.WaitGroup{},
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

	s.userWG.Add(1)
	// Launch the goroutine to save data concurrently
	go func() {
		defer s.userWG.Done()
		err := s.userRepo.CreateUser(user)
		if err != nil {
			fmt.Println("Error creating user: ", err)
		}
	}()

	// set the user as the active user
	globals.ActiveUser = user.Email
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

	// set the user as the active user
	globals.ActiveUser = user.Email
	return &user, nil
}

func (s *UserService) GetAllUsers() []entities.User {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil
	}
	return users
}

func (s *UserService) AddWinToUser(user *entities.User, gameId string) error {
	return s.userRepo.AddWin(user.UserId, gameId)
}

func (s *UserService) AddLossToUser(user *entities.User, gameId string) error {
	return s.userRepo.AddLoss(user.UserId, gameId)
}

func (s *UserService) WaitForCompletion() {
	s.userWG.Wait()
}
