package services

import (
	"project2/internal/domain/interfaces"
	"sync"
)

type LeaderboardService struct {
	leaderBoardRepo interfaces.LeaderboardRepository
	userService     *UserService
	leaderboardWG   *sync.WaitGroup
}

func NewLeaderboardService(leaderBoardRepo interfaces.LeaderboardRepository, userService *UserService) *LeaderboardService {
	return &LeaderboardService{
		leaderBoardRepo: leaderBoardRepo,
		userService:     userService,
		leaderboardWG:   &sync.WaitGroup{},
	}
}

//func (s *LeaderboardService) GetOverallLeaderboard() []entities.User {
//	users := s.userService.GetAllUsers()
//	if len(users) == 0 {
//		return nil
//	}
//	overAllLeaderboard, err := s.leaderBoardRepo.GetOverallLeaderboard(users)
//	if err != nil {
//		return nil
//	}
//	return overAllLeaderboard
//
//}
//
//func (s *LeaderboardService) GetGameLeaderboard(gameId string) []entities.User {
//	users := s.userService.GetAllUsers()
//	if len(users) == 0 {
//		return nil
//	}
//	gameLeaderboard, err := s.leaderBoardRepo.GetGameLeaderboard(gameId, users)
//	if err != nil {
//		return nil
//	}
//	return gameLeaderboard
//}
