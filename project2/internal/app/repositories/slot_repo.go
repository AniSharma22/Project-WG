package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"reflect"
	"time"
)

func init() {
	if _, err := os.Stat(config.SlotsFile); err == nil {
		go loadAllResults()
	}
}

type slotRepo struct {
	slots []entities.Slot // This can be replaced with actual database or file storage
}

func NewSlotRepo() interfaces.SlotRepository {
	return &slotRepo{}
}

func (r *slotRepo) GetSlotsByDate(date string, gameId string) ([]entities.SlotStats, error) {
	slots, exists := globals.SlotsMap[date][gameId]
	if !exists {
		return nil, errors.New("today's Date slots are not available")
	}
	return slots, nil
}

func (r *slotRepo) GetSlotByDateAndTime(date string, gameId string, time time.Time) (entities.SlotStats, error) {
	for _, v := range globals.SlotsMap[date][gameId] {
		if v.Time == time {
			return v, nil
		}
	}
	return entities.SlotStats{}, errors.New("this slot data is not available")
}

func (r *slotRepo) BookSlot(date string, gameId string, time time.Time, user entities.User) error {










		// Access the slice of slots for the given date and gameId
	slots := globals.SlotsMap[date][gameId]

	// Find the slot with the specified time and update the BookedBy field
	for i, slot := range slots {
		if slot.Time == time {
			// Update the BookedBy slice
			slot.BookedBy = append(slot.BookedBy, user)

			// Reassign the updated slot back to the slice
			slots[i] = slot

			// Update the map with the modified slice
			globals.SlotsMap[date][gameId] = slots

			return nil // Return early if the slot was successfully booked
		}
	}



	file, err := os.OpenFile(config.SlotsFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var slots2 [] entities.Slot

	// Decode existing results from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&slots2); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing slots:", err)
		return err
	}

	for _, v := range slots2 {
		if v.Date == date{

		}
	}

	// Truncate the file to overwrite it with the updated results array
	if err := file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}

	// Move the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return err
	}

	// Encode the updated games array to JSON and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(slots2); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err























	return fmt.Errorf("slot not found or already booked")
}

func (r *slotRepo) GetPendingInvites(userID string) ([]entities.Slot, error) {
	// Filter and return slots with pending invites for the user
	var pendingInvites []entities.Slot
	for _, slot := range r.slots {
		for _, invitedUser := range slot.InvitedUsers {
			if invitedUser.UserId == userID {
				pendingInvites = append(pendingInvites, slot)
				break
			}
		}
	}
	return pendingInvites, nil
}

func (r *slotRepo) UpdateSlot(slot entities.Slot) error {
	// Update the slot in the storage
	for i, s := range r.slots {
		if s.SlotID == slot.SlotID {
			r.slots[i] = slot
			return nil
		}
	}
	return errors.New("slot not found")
}

func loadAllSlots() {
	slotDataChan := make(chan any)
	go utils.StreamJSONObjects(slotDataChan, config.SlotsFile, reflect.TypeOf(entities.Slot{})) // Corrected type here

	for slot := range slotDataChan {
		slotData, ok := slot.(*entities.Slot)
		if !ok {
			fmt.Println("Error: received data is not of type *entities.Slot")
			continue
		}
		// Initialize the date map if not already present
		if _, exists := globals.SlotsMap[slotData.Date]; !exists {
			globals.SlotsMap[slotData.Date] = make(map[string][]entities.SlotStats)
		}
		// Insert game slots into the appropriate date and gameID map
		for _, v := range slotData.Games {
			globals.SlotsMap[slotData.Date][v.GameID] = v.Slots
		}
	}

	// Allow some time for the slot data to load
	time.Sleep(2 * time.Second)
	dateStr := time.Now().Format("2006-01-02") // Format date as "YYYY-MM-DD"

	// Check if today's entry exists
	_, exists := globals.SlotsMap[dateStr]
	if !exists {
		// Initialize map for today's date
		globals.SlotsMap[dateStr] = make(map[string][]entities.SlotStats)

		// Iterate over all game IDs
		for gameId := range globals.GamesMap {
			var slots []entities.SlotStats

			// Generate slots from 09:00 AM to 06:00 PM with 20-minute intervals
			startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local)
			endTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 0, 0, 0, time.Local)

			for startTime.Before(endTime) {
				slotId, _ := utils.GetUuid() // Generate slot ID
				slot := entities.SlotStats{
					SlotID:       slotId,
					Time:         startTime,
					BookedBy:     make([]entities.User, 0),
					InvitedUsers: make([]entities.User, 0),
					Duration:     20 * time.Minute,
					IsBooked:     false,
				}
				slots = append(slots, slot)
				startTime = startTime.Add(20 * time.Minute)
			}

			// Add the generated slots for the current game to the map
			globals.SlotsMap[dateStr][gameId] = slots
		}
	}
}



}