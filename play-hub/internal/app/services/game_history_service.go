package services

import (
	"fmt"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameHistoryService struct {
	gameHistoryRepo interfaces.GameHistoryRepository
	userService     interfaces.UserService
	SlotService     interfaces.SlotService
	gameHistoryWG   *sync.WaitGroup
}

func NewGameHistoryService(gameHistoryRepo interfaces.GameHistoryRepository, userService interfaces.UserService, slotService interfaces.SlotService) interfaces.GameHistoryService {
	return &GameHistoryService{
		gameHistoryRepo: gameHistoryRepo,
		userService:     userService,
		SlotService:     slotService,
		gameHistoryWG:   &sync.WaitGroup{},
	}
}

func (gh *GameHistoryService) GetTotalGameHistory() ([]entities.GameHistory, error) {
	user, err := gh.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return nil, err
	}

	return gh.gameHistoryRepo.GetUserGameHistory(user.ID)
}

func (gh *GameHistoryService) GetResultsToUpdate() ([]entities.GameHistory, error) {
	// Retrieve the active user by email
	user, err := gh.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user by email: %w", err)
	}

	// Get the current time
	now := time.Now()

	// Fetch game histories that need updating for the user
	gameHistories, err := gh.gameHistoryRepo.GetResultsToUpdate(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error fetching game histories: %w", err)
	}

	// Initialize a slice to hold game histories that need updates
	var resultsToUpdate []entities.GameHistory

	// Iterate over the fetched game histories
	for _, gameHistory := range gameHistories {
		// Retrieve the slot associated with the game history
		slot, err := gh.SlotService.GetSlotById(gameHistory.SlotID)
		if err != nil {
			// Log the error but continue processing other game histories
			continue
		}

		// Check if the current time is before the slot's end time
		if now.After(slot.EndTime) {
			resultsToUpdate = append(resultsToUpdate, gameHistory)
		}
	}

	// Return the list of game histories that need updates
	return resultsToUpdate, nil
}

func (gh *GameHistoryService) UpdateResult(result string, slotId primitive.ObjectID) error {
	user, err := gh.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return err
	}
	err = gh.gameHistoryRepo.UpdateResult(result, slotId, user.ID)
	if err != nil {
		return err
	}
	err = gh.userService.AddResult(user.ID, result)
	if err != nil {
		return err
	}
	err = gh.SlotService.AddResultToSlot(user.ID, slotId, result)
	if err != nil {
		return err
	}
	return nil
}
