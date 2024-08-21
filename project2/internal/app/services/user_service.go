package services

import (
	"errors"
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

	user.Role = "user"
	// Hash the password (assuming a method for hashing exists)
	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

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

func (s *UserService) Login(email string, password string) (*entities.User, error) {
	hashedPassword, _ := utils.GetHashedPassword(password)
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != hashedPassword {
		return nil, errors.New("wrong password")
	}
	globals.ActiveUser = user.Email
	return user, nil
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
