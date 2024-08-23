package services

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"sort"
	"time"
)

type SlotService struct {
	slotRepo    interfaces.SlotRepository
	UserService *UserService
}

func NewSlotService(slotRepo interfaces.SlotRepository, userService *UserService) *SlotService {
	return &SlotService{
		slotRepo:    slotRepo,
		UserService: userService,
	}
}

// Helper function to parse time strings into time.Time
func parseSlotTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04", timeStr)
}

func (s *SlotService) GetGameSlots(game *entities.Game) ([]entities.Slot, error) {
	todayDate := time.Now().Format("2006-01-02")

	// Fetch slots for the given game on today's date
	slots, err := s.slotRepo.GetSlotsByDate(todayDate, game.ID)
	if err != nil {
		return nil, err
	}

	// Sort slots by StartTime
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].StartTime < slots[j].StartTime
	})

	return slots, nil
}

func (s *SlotService) BookSlot(game *entities.Game, slot *entities.Slot) error {
	// Parse the slot start time
	slotStartTime, err := parseSlotTime(slot.StartTime)
	if err != nil {
		return fmt.Errorf("invalid slot start time format: %v", err)
	}

	// Capture today's date for comparison
	now := time.Now()

	// Combine today's date with the slot start time
	slotStartTime = time.Date(
		now.Year(), now.Month(), now.Day(),
		slotStartTime.Hour(), slotStartTime.Minute(), 0, 0,
		time.Local,
	)

	// Check if the slot timing has passed
	if now.After(slotStartTime) {
		return fmt.Errorf("cannot book slot: the slot timing has already passed")
	}

	// Check if the slot already has 4 members booked
	if len(slot.BookedUsers) >= 4 {
		return fmt.Errorf("cannot book slot: the slot is already fully booked")
	}

	// Proceed to book the slot using the repository
	user, err := s.UserService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}

	err = s.slotRepo.BookSlot(user.ID, now.Format("2006-01-02"), game.ID, slot.StartTime)
	if err != nil {
		return fmt.Errorf("failed to book slot: %v", err)
	}

	return nil
}

func (s *SlotService) InviteToSlot(userId primitive.ObjectID, game *entities.Game, slot *entities.Slot) error {
	// Convert slot.StartTime to time.Time
	slotStartTime, err := parseSlotTime(slot.StartTime)
	if err != nil {
		return fmt.Errorf("invalid slot start time format: %v", err)
	}

	// Capture today's date for comparison
	now := time.Now()

	// Combine today's date with the slot start time
	slotStartTime = time.Date(
		now.Year(), now.Month(), now.Day(),
		slotStartTime.Hour(), slotStartTime.Minute(), 0, 0,
		time.Local,
	)

	// Check if the slot timing has passed
	if now.After(slotStartTime) {
		return fmt.Errorf("cannot invite: the slot timing has already passed")
	}

	// Check if the slot already has 4 members booked
	if len(slot.BookedUsers) >= 4 {
		return fmt.Errorf("cannot invite: the slot is already fully booked")
	}

	// Call the repository's InviteToSlot method with the correct parameters
	err = s.slotRepo.InviteToSlot(userId, now.Format("2006-01-02"), game.ID, slot.StartTime)
	if err != nil {
		return fmt.Errorf("failed to invite user: %v", err)
	}

	return nil
}

func (s *SlotService) GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error) {
	return s.slotRepo.GetSlotById(slotId)
}
