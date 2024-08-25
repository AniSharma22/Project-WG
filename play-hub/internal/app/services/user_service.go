package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/internal/models"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"sync"
	"time"
)

type UserService struct {
	userRepo    interfaces.UserRepository
	SlotService *SlotService
	GameService *GameService
	userWG      *sync.WaitGroup
}

func NewUserService(userRepo interfaces.UserRepository, slotService *SlotService, gameService *GameService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		SlotService: slotService,
		GameService: gameService,
		userWG:      &sync.WaitGroup{},
	}
}

func (s *UserService) Signup(user *entities.User) error {

	user.Role = "user"
	user.InvitedSlots = []primitive.ObjectID{}

	// Create the user
	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	// Signup successful
	globals.ActiveUser = user.Email
	return nil
}

func (s *UserService) EmailAlreadyExists(email string) bool {
	err := s.userRepo.EmailAlreadyExists(email)
	if err != nil {
		return false
	}
	return true
}

func (s *UserService) Login(email string, password []byte) (*entities.User, error) {

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("wrong password")
	}
	globals.ActiveUser = user.Email
	return user, nil
}

func (s *UserService) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
	return s.userRepo.GetUserById(userId)
}

func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// returns all the pending invites for upcoming games and deletes the expired ones
func (s *UserService) GetPendingInvites() ([]models.Invite, error) {
	// Retrieve all pending invites
	invites, err := s.userRepo.GetPendingInvites(globals.ActiveUser)
	if err != nil {
		return nil, err
	}

	// Current time
	now := time.Now()

	// List to hold valid (non-expired) invites
	var validInvites []models.Invite

	// Iterate over invites and filter out expired ones
	for _, invite := range invites {
		slot, _ := s.SlotService.GetSlotById(invite)
		if slot.StartTime.Before(now) {
			// Remove the expired invite from the repository
			err := s.userRepo.DeleteInvite(invite)
			if err != nil {
				return nil, err
			}
		} else {
			// Add valid invite to the list by creating a invite object
			inviteToAdd := models.Invite{
				SlotId: slot.ID,
				GameName: func() string {
					game, _ := s.GameService.GetGameByID(slot.GameID)
					return game.Name
				}(),
				Date:      slot.Date,
				StartTime: slot.StartTime,
				EndTime:   slot.EndTime,
				BookedUsers: func() []string {
					var bookedUsers []string
					for _, userId := range slot.BookedUsers {
						user, _ := s.userRepo.GetUserById(userId)
						bookedUsers = append(bookedUsers, utils.GetNameFromEmail(user.Email))
					}
					return bookedUsers
				}(),
			}
			validInvites = append(validInvites, inviteToAdd)
		}
	}

	return validInvites, nil
}

func (s *UserService) AcceptInvite(slotId primitive.ObjectID) error {
	slot, err := s.SlotService.GetSlotById(slotId)
	if err != nil {
		return err
	}
	game, err := s.GameService.GetGameByID(slot.GameID)
	if err != nil {
		return err
	}

	err = s.SlotService.BookSlot(game, slot)
	if err != nil {
		return err
	}

	// Delete the invite from the user's invitedSlots
	err = s.userRepo.DeleteInvite(slotId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) RejectInvite(slotId primitive.ObjectID) error {
	err := s.userRepo.DeleteInvite(slotId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) AddResult(userId primitive.ObjectID, result string) error {
	if result == "win" {
		s.userRepo.AddWin(userId)
	} else if result == "loss" {
		s.userRepo.AddLoss(userId)
	} else {
		return errors.New("wrong result")
	}
	return nil
}

func (s *UserService) GetAllUsersByScore() ([]entities.User, error) {
	return s.userRepo.GetAllUsersByScore()
}
