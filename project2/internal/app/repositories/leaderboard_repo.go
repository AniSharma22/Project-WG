package repositories

import (
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"sort"
)

type leaderboardRepo struct {
}

func NewLeaderboardRepo() interfaces.Leaderboard {
	return &leaderboardRepo{}
}

func (r *leaderboardRepo) GetGameLeaderboard(gameId string, users []entities.User) ([]entities.User, error) {
	var gameIndex int
	for i, game := range users[0].GameStats {
		if game.GameID == gameId {
			gameIndex = i
		}
	}
	sortByGameScore(gameIndex, users)
	return users, nil
}

func (r *leaderboardRepo) GetOverallLeaderboard(users []entities.User) ([]entities.User, error) {
	sortByScore(users)
	return users, nil
}

func sortByScore(users []entities.User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})
}

func sortByGameScore(gameIndex int, users []entities.User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].GameStats[gameIndex].Score > users[j].GameStats[gameIndex].Score
	})
}
