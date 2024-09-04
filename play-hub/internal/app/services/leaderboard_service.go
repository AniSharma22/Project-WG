package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"project2/internal/models"
	"project2/pkg/utils"
	"sync"
)

type LeaderboardService struct {
	leaderBoardRepo repository_interfaces.LeaderboardRepository
	bookingService  service_interfaces.BookingService
	leaderboardWG   *sync.WaitGroup
}

func NewLeaderboardService(leaderBoardRepo repository_interfaces.LeaderboardRepository, bookingService service_interfaces.BookingService) service_interfaces.LeaderboardService {
	return &LeaderboardService{
		leaderBoardRepo: leaderBoardRepo,
		bookingService:  bookingService,
		leaderboardWG:   &sync.WaitGroup{},
	}
}

func (s *LeaderboardService) GetGameLeaderboard(ctx context.Context, gameId uuid.UUID) ([]models.Leaderboard, error) {
	return s.leaderBoardRepo.FetchGameLeaderboard(ctx, gameId)
}

func (s *LeaderboardService) AddWinToUser(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, bookingId uuid.UUID) error {
	userStats, err := s.leaderBoardRepo.FetchUserGameStats(ctx, userId, gameId)
	if err != nil {
		return fmt.Errorf("failed to fetch user stats for game %s: %w", gameId, err)
	}
	userStats.Wins++
	newUserScore := utils.GetTotalScore(userStats.Wins, userStats.Losses)
	userStats.Score = float64(newUserScore)
	err = s.leaderBoardRepo.UpdateUserGameStats(ctx, userStats)
	if err != nil {
		return fmt.Errorf("failed to update user stats for game %s: %w", gameId, err)
	}
	err = s.bookingService.UpdateBookingResult(ctx, bookingId, "win")
	if err != nil {
		return fmt.Errorf("failed to update booking result for game %s: %w", gameId, err)
	}
	return nil

}

func (s *LeaderboardService) AddLossToUser(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, bookingId uuid.UUID) error {
	userStats, err := s.leaderBoardRepo.FetchUserGameStats(ctx, userId, gameId)
	if err != nil {
		return fmt.Errorf("failed to fetch user stats for game %s: %w", gameId, err)
	}
	userStats.Losses++
	newUserScore := utils.GetTotalScore(userStats.Wins, userStats.Losses)
	userStats.Score = float64(newUserScore)
	err = s.leaderBoardRepo.UpdateUserGameStats(ctx, userStats)
	if err != nil {
		return fmt.Errorf("failed to update user stats for game %s: %w", gameId, err)
	}
	err = s.bookingService.UpdateBookingResult(ctx, bookingId, "loss")
	if err != nil {
		return fmt.Errorf("failed to update booking result for game %s: %w", gameId, err)
	}
	return nil
}
