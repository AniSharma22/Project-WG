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
		return false
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

//func (s *UserService) Signup(user *entities.User) error {
//
//	user.Role = "user"
//	user.InvitedSlots = []primitive.ObjectID{}
//
//	// Create the user
//	if err := s.userRepo.CreateUser(user); err != nil {
//		return err
//	}
//
//	// Signup successful
//	globals.ActiveUser = user.Email
//	return nil
//}
//
//func (s *UserService) EmailAlreadyExists(email string) bool {
//	err := s.userRepo.EmailAlreadyExists(email)
//	if err != nil {
//		return false
//	}
//	return true
//}
//
//func (s *UserService) Login(email string, password []byte) (*entities.User, error) {
//
//	user, err := s.userRepo.GetUserByEmail(email)
//	if err != nil {
//		return nil, err
//	}
//	if !utils.VerifyPassword(password, user.Password) {
//		return nil, errors.New("wrong password")
//	}
//	globals.ActiveUser = user.Email
//	return user, nil
//}
//
//func (s *UserService) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
//	return s.userRepo.GetUserById(userId)
//}
//
//func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
//	return s.userRepo.GetUserByEmail(email)
//}
//
//// returns all the pending invites for upcoming games and deletes the expired ones
//func (s *UserService) GetPendingInvites() ([]models.Invite, error) {
//	// Retrieve all pending invites
//	invites, err := s.userRepo.GetPendingInvites(globals.ActiveUser)
//	if err != nil {
//		return nil, err
//	}
//
//	// Current time
//	now := time.Now()
//
//	// List to hold valid (non-expired) invites
//	var validInvites []models.Invite
//
//	// Iterate over invites and filter out expired ones
//	for _, invite := range invites {
//		slot, _ := s.SlotService.GetSlotById(invite)
//		if slot.StartTime.Before(now) {
//			// Remove the expired invite from the repository
//			err := s.userRepo.DeleteInvite(invite)
//			if err != nil {
//				return nil, err
//			}
//		} else {
//			// Add valid invite to the list by creating an invitation object
//			inviteToAdd := models.Invite{
//				SlotId: slot.ID,
//				GameName: func() string {
//					game, _ := s.GameService.GetGameByID(slot.GameID)
//					return game.Name
//				}(),
//				Date:      slot.Date,
//				StartTime: slot.StartTime,
//				EndTime:   slot.EndTime,
//				BookedUsers: func() []string {
//					var bookedUsers []string
//					for _, userId := range slot.BookedUsers {
//						user, _ := s.userRepo.GetUserById(userId)
//						bookedUsers = append(bookedUsers, utils.GetNameFromEmail(user.Email))
//					}
//					return bookedUsers
//				}(),
//			}
//			validInvites = append(validInvites, inviteToAdd)
//		}
//	}
//
//	return validInvites, nil
//}
//
//func (s *UserService) AcceptInvite(slotId primitive.ObjectID) error {
//	slot, err := s.SlotService.GetSlotById(slotId)
//	if err != nil {
//		return err
//	}
//
//	game, err := s.GameService.GetGameByID(slot.GameID)
//	if err != nil {
//		return err
//	}
//
//	err = s.SlotService.BookSlot(game, slot)
//	if err != nil {
//		return err
//	}
//
//	// Delete the invite from the user's invitedSlots
//	err = s.userRepo.DeleteInvite(slotId)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *UserService) RejectInvite(slotId primitive.ObjectID) error {
//	err := s.userRepo.DeleteInvite(slotId)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *UserService) AddResult(userId primitive.ObjectID, result string) error {
//	if result == "win" {
//		s.userRepo.AddWin(userId)
//	} else if result == "loss" {
//		s.userRepo.AddLoss(userId)
//	} else {
//		return errors.New("wrong result")
//	}
//	return nil
//}
//
//func (s *UserService) GetAllUsersByScore() ([]entities.User, error) {
//	return s.userRepo.GetAllUsersByScore()
//}
