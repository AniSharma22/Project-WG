package utils

import (
	"github.com/google/uuid"
	"math"
)

// GetUuid generates a new random UUID and returns it as a string
func GetUuid() (string, error) {
	// Generate a new UUID
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func GetTotalScore(totalWins, totalLosses int) float32 {
	totalGames := totalWins + totalLosses
	return calculateScore(totalWins, totalLosses, totalGames)
}

func GetGameScore(totalWins, totalLosses, totalGames int) float32 {
	return calculateScore(totalWins, totalLosses, totalGames)
}

func calculateScore(totalWins, totalLosses, totalGames int) float32 {
	var winLossRatio float32 = float32(totalWins) / float32(totalLosses)
	var gameFactor float32 = float32(1) + float32(math.Sqrt(float64(totalGames)))
	return (winLossRatio * gameFactor) / 100
}
