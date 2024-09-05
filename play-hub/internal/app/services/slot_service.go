package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"sync"
	"time"
)

type SlotService struct {
	slotRepo repository_interfaces.SlotRepository
	userWG   *sync.WaitGroup
}

func NewSlotService(slotRepo repository_interfaces.SlotRepository) service_interfaces.SlotService {
	return &SlotService{
		slotRepo: slotRepo,
		userWG:   &sync.WaitGroup{},
	}
}

// GetCurrentDayGameSlots retrieves all slots for the current day for a specific game.
func (s *SlotService) GetCurrentDayGameSlots(ctx context.Context, gameID uuid.UUID) ([]entities.Slot, error) {

	// Call the repository to fetch slots by game ID and date
	slots, err := s.slotRepo.FetchSlotsByGameIDAndDate(ctx, gameID, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch slots for game ID %s on date %s: %w", gameID, time.Now().Format("2006-03-04"), err)
	}

	return slots, nil
}
func (s *SlotService) GetSlotByID(ctx context.Context, slotID uuid.UUID) (*entities.Slot, error) {
	return s.slotRepo.FetchSlotByID(ctx, slotID)
}

func (s *SlotService) MarkSlotAsBooked(ctx context.Context, slotID uuid.UUID) error {
	return s.slotRepo.UpdateSlotStatus(ctx, slotID, true)
}
