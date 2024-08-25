package services

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"slices"
	"sort"
	"time"
)

type SlotService struct {
	slotRepo        interfaces.SlotRepository
	userRepo        interfaces.UserRepository
	gameHistoryRepo interfaces.GameHistoryRepository
}

func NewSlotService(slotRepo interfaces.SlotRepository, userRepo interfaces.UserRepository, gameHistoryRepo interfaces.GameHistoryRepository) *SlotService {
	return &SlotService{
		slotRepo:        slotRepo,
		userRepo:        userRepo,
		gameHistoryRepo: gameHistoryRepo,
	}
}

// Helper function to parse time strings into time.Time
func parseSlotTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04", timeStr)
}

// Returns all the game slots for today for the given game
func (s *SlotService) GetGameSlots(game *entities.Game) ([]entities.Slot, error) {
	todayDate := time.Now().Truncate(24 * time.Hour)

	// Fetch slots for the given game on today's date
	slots, err := s.slotRepo.GetSlotsByDate(todayDate, game.ID)
	if err != nil {
		return nil, err
	}

	// Sort slots by StartTime
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].StartTime.Before(slots[j].StartTime)
	})

	return slots, nil
}

// Books a slot for the user for the given corresponding game
func (s *SlotService) BookSlot(game *entities.Game, slot *entities.Slot) error {
	now := time.Now()

	// Check if the slot timing has passed
	if now.After(slot.StartTime) {
		return fmt.Errorf("cannot book slot: the slot timing has already passed")
	}

	// Check if the slot already has 4 members booked
	if len(slot.BookedUsers) >= 4 {
		return fmt.Errorf("cannot book slot: the slot is already fully booked")
	}

	// fetch the user active user details
	user, err := s.userRepo.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}

	// check if the user has already booked this slot
	if slices.Contains(slot.BookedUsers, user.ID) {
		return fmt.Errorf("you have already booked this slot before")
	}
	// Proceed to book the slot using the repository
	err = s.slotRepo.BookSlot(user.ID, slot.ID)
	if err != nil {
		return fmt.Errorf("failed to book slot: %v", err)
	}
	history := entities.GameHistory{
		ID:        primitive.NewObjectID(),
		UserID:    user.ID,
		GameID:    game.ID,
		SlotID:    slot.ID,
		Result:    "",
		CreatedAt: time.Now(),
	}
	// Also add the game booking to game history table
	s.gameHistoryRepo.AddGameHistory(&history)

	return nil
}

func (s *SlotService) InviteToSlot(userId primitive.ObjectID, game *entities.Game, slot *entities.Slot) error {
	// Capture today's date for comparison
	now := time.Now()

	// Check if the slot timing has passed
	if now.After(slot.StartTime) {
		return fmt.Errorf("cannot invite: the slot timing has already passed")
	}

	// Check if the slot already has 4 members booked
	if len(slot.BookedUsers) >= 4 {
		return fmt.Errorf("cannot invite: the slot is already fully booked")
	}

	// Check if the user has already booked in this slot
	if slices.Contains(slot.BookedUsers, userId) {
		return fmt.Errorf("cannot invite: the user has already booked in this slot")
	}

	// send invite to the concerned user
	err := s.userRepo.AddToInvites(userId, slot.ID)
	if err != nil {
		return fmt.Errorf("failed to invite: %v", err)
	}
	return nil
}

func (s *SlotService) GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error) {
	slot, err := s.slotRepo.GetSlotById(slotId)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (s *SlotService) GetUpcomingBookedSlots() ([]entities.Slot, error) {
	user, _ := s.userRepo.GetUserByEmail(globals.ActiveUser)
	return s.slotRepo.GetUpcomingBookedSlots(user.ID)
}

func (s *SlotService) AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error {
	return s.slotRepo.AddResultToSlot(userId, slotId, result)
}
