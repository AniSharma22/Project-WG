package services

import (
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"sync"
)

type LeaderboardService struct {
	leaderBoardRepo interfaces.LeaderboardRepository
	userService     interfaces.UserService
	leaderboardWG   *sync.WaitGroup
}

func NewLeaderboardService(leaderBoardRepo interfaces.LeaderboardRepository, userService interfaces.UserService) interfaces.LeaderboardService {
	return &LeaderboardService{
		leaderBoardRepo: leaderBoardRepo,
		userService:     userService,
		leaderboardWG:   &sync.WaitGroup{},
	}
}

func (s *LeaderboardService) GetOverallLeaderboard() ([]entities.User, error) {
	return s.userService.GetAllUsersByScore()
}
