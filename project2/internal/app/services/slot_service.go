package services

import (
	"project2/internal/domain/interfaces"
)

type SlotService struct {
	slotRepo interfaces.SlotRepository
}

func NewSlotService(slotRepo interfaces.SlotRepository) *SlotService {
	return &SlotService{
		slotRepo: slotRepo,
	}
}

//func (s *SlotService) GetSlotsForToday() ([]entities.SlotStats, error) {
//	return s.slotRepo.GetSlotsByDate(time.DateOnly, "1")
//}
//
////func (s *SlotService) GetSlotsForToday() ([]entities.Slot, error) {
////	today := time.Now()
////	slots, err := s.slotRepo.GetSlotsByDate(today)
////	if err != nil {
////		return nil, err
////	}
////	return slots, nil
////}
////
////func (s *SlotService) BookSlot(slotID string, userID string) error {
////	// Retrieve user details
////	user, exists := globals.UsersMap[userID]
////	if !exists {
////		return errors.New("user not found")
////	}
////
////	// Book the slot
////	err := s.slotRepo.BookSlot(slotID, user)
////	if err != nil {
////		return err
////	}
////	fmt.Println("Slot booked successfully.")
////	return nil
////}
////
////func (s *SlotService) GetPendingInvites(userID string) ([]entities.Slot, error) {
////	return s.slotRepo.GetPendingInvites(userID)
////}
////
////func (s *SlotService) ExpireOldInvites() error {
////	// Get current time and expire invites for past slots
////	for _, slot := range globals.SlotsMap {
////		if slot.Time.Before(time.Now()) {
////			// Expire invites for this slot
////			slot.InvitedUsers = []entities.User{}
////			if err := s.slotRepo.UpdateSlot(slot); err != nil {
////				return err
////			}
////		}
////	}
////	return nil
////}
