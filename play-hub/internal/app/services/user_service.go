package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"sync"
)

type UserService struct {
	userRepo repository_interfaces.UserRepository
	userWG   *sync.WaitGroup
}

func NewUserService(userRepo repository_interfaces.UserRepository) service_interfaces.UserService {
	return &UserService{
		userRepo: userRepo,
		userWG:   &sync.WaitGroup{},
	}
}

// Signup registers a new user in the system.
func (s *UserService) Signup(ctx context.Context, user *entities.User) error {
	// Check if email is already registered
	exists := s.EmailAlreadyRegistered(ctx, user.Email)

	if exists {
		return errors.New("email already registered")
	}
	user.Username = utils.GetNameFromEmail(user.Email)

	// Create user
	userId, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	globals.ActiveUser = userId
	return nil
}

// Login authenticates a user with email and password.
func (s *UserService) Login(ctx context.Context, email string, password []byte) (*entities.User, error) {
	// Fetch the user by email
	user, err := s.userRepo.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Compare the stored hashed password with the provided password
	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	globals.ActiveUser = user.UserID
	return user, nil
}

// EmailAlreadyRegistered checks if an email is already registered in the system.
func (s *UserService) EmailAlreadyRegistered(ctx context.Context, email string) bool {
	user, err := s.userRepo.FetchUserByEmail(ctx, email)
	if err != nil {
		return true
	}
	return user != nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	return s.userRepo.FetchUserById(ctx, userID)
}

// GetUserByEmail retrieves a user by their email address.
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return s.userRepo.FetchUserByEmail(ctx, email)
}

// GetUserByUsername retrieves a user by their username.
func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	return s.userRepo.FetchUserByUsername(ctx, username)
}
