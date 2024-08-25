package utils

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"time"
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

func GetNameFromEmail(email string) string {
	var name bytes.Buffer
	for i := 0; i < len(email); i++ {
		if email[i] == '.' {
			name.WriteByte(' ')
		} else if email[i] == '@' {
			break
		} else {
			name.WriteByte(email[i])
		}
	}
	return name.String()
}

func calculateScore(totalWins, totalLosses, totalGames int) float32 {
	var winLossRatio float32 = float32(totalWins) / float32(totalLosses)
	var gameFactor float32 = float32(1) + float32(math.Sqrt(float64(totalGames)))
	return (winLossRatio * gameFactor) / 100
}

func InsertAllSlots(slotRepo interfaces.SlotRepository, gameRepo interfaces.GameRepository) error {
	today := time.Now().Truncate(24 * time.Hour)

	// Fetch all games
	games, err := gameRepo.GetAllGames()
	if err != nil {
		return fmt.Errorf("error fetching games: %w", err)
	}

	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)

	for _, game := range games {
		// Check for existing slots for this game on today's date
		existingSlots, err := slotRepo.GetSlotsByDate(today, game.ID)
		if err != nil {
			return fmt.Errorf("error checking existing slots for game %s: %w", game.Name, err)
		}

		// If no slots exist, create new slots
		if len(existingSlots) == 0 {
			for current := startTime; current.Before(endTime); current = current.Add(20 * time.Minute) {
				slotEndTime := current.Add(20 * time.Minute)
				if slotEndTime.After(endTime) {
					slotEndTime = endTime
				}

				newSlot := entities.Slot{
					ID:          primitive.NewObjectID(),
					GameID:      game.ID,
					Date:        time.Now().Truncate(24 * time.Hour),
					StartTime:   current,
					EndTime:     slotEndTime,
					BookedUsers: []primitive.ObjectID{},
					Results:     []entities.Result{},
				}

				// Insert the new slot
				if _, err := slotRepo.InsertSlot(newSlot); err != nil {
					return fmt.Errorf("error inserting slot for game %s: %w", game.Name, err)
				}
			}
		}
	}

	return nil
}
