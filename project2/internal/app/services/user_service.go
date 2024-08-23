package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
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

// GetPendingInvites retrieves pending invites and removes expired ones
func (s *UserService) GetPendingInvites(email string) ([]entities.InvitedSlot, error) {
	// Retrieve all pending invites
	invites, err := s.userRepo.GetPendingInvites(email)
	if err != nil {
		return nil, err
	}

	// Current time
	now := time.Now()

	// Iterate over invites and remove expired ones
	for _, invite := range invites {
		startTime, _ := parseSlotTime(invite.StartTime)
		if startTime.Before(now) {
			// Remove the expired invite from the repository
			err := s.userRepo.DeleteInvite(invite.SlotID)
			if err != nil {
				return nil, err
			}
		}
	}

	// Retrieve updated list of pending invites
	updatedInvites, err := s.userRepo.GetPendingInvites(email)
	if err != nil {
		return nil, err
	}

	return updatedInvites, nil
}

func (s *UserService) AcceptInvite(slotId primitive.ObjectID) error {
	// Retrieve the user
	user, err := s.userRepo.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return err
	}

	// Retrieve the slot
	slot, err := s.SlotService.GetSlotById(slotId)
	if err != nil {
		return err
	}

	// Check if there is space in the slot
	if len(slot.BookedUsers) >= slot.GameID {
		return errors.New("no space available in the slot")
	}

	// Add user to the slot's bookedUsers array
	err = s.slotRepo.AddUserToSlot(slotId, user.ID)
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

}

//func (s *UserService) Signup(user *entities.User) error {
//	// Set user role
//	user.Role = "user"
//	//Initialize GameStats with all GameIDs from the global map
//	user.GameStats = make([]entities.GameStats, 0, len(globals.GamesMap))
//	for gameID := range globals.GamesMap {
//		user.GameStats = append(user.GameStats, entities.GameStats{
//			GameID:     gameID,
//			Wins:       0,
//			Losses:     0,
//			TotalGames: 0,
//			Score:      0.0,
//		})
//	}
//
//	// Generate UUID for the user
//	uuid, err := utils.GetUuid()
//	if err != nil {
//		fmt.Println("Error generating UUID: ", err)
//		return err
//	}
//	user.UserId = uuid
//
//	// Hash the user's password
//	hashedPassword, err := utils.GetHashedPassword(user.Password)
//	if err != nil {
//		fmt.Println("Error generating password hash: ", err)
//		return err
//	}
//	user.Password = hashedPassword
//
//	s.userWG.Add(1)
//	// Launch the goroutine to save data concurrently
//	go func() {
//		defer s.userWG.Done()
//		err := s.userRepo.CreateUser(user)
//		if err != nil {
//			fmt.Println("Error creating user: ", err)
//		}
//	}()
//
//	// set the user as the active user
//	globals.ActiveUser = user.Email
//	return nil
//}
//
//func (s *UserService) Login(email string, password string) (*entities.User, error) {
//	user, exists := globals.UsersMap[email]
//	if !exists {
//		return nil, errors.New("user not found")
//	}
//
//	isValidUser := utils.VerifyPassword(user.Password, password)
//	if !isValidUser {
//		return nil, errors.New("wrong password")
//	}
//
//	// set the user as the active user
//	globals.ActiveUser = user.Email
//	return &user, nil
//}
//
//func (s *UserService) GetAllUsers() []entities.User {
//	users, err := s.userRepo.GetAllUsers()
//	if err != nil {
//		return nil
//	}
//	return users
//}
//
//func (s *UserService) AddWinToUser(user *entities.User, gameId string) error {
//	return s.userRepo.AddWin(user.UserId, gameId)
//}
//
//func (s *UserService) AddLossToUser(user *entities.User, gameId string) error {
//	return s.userRepo.AddLoss(user.UserId, gameId)
//}
//
//func (s *UserService) WaitForCompletion() {
//	s.userWG.Wait()
//}
