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
	"time"
)

func init() {
	if _, err := os.Stat(config.SlotsFile); err == nil {
		go loadAllSlots()
	}
}

type slotRepo struct {
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

func (r *slotRepo) BookSlot(user entities.User, date string, gameId string, time time.Time) error {
	fetchedSlot, _ := r.GetSlotsByDate(date, gameId)
	for _, slots := range fetchedSlot {
		if slots.Time == time {
			slots.BookedBy = append(slots.BookedBy, user)
		}
	}

	globals.SlotsMap[date][gameId] = fetchedSlot

	file, err := os.OpenFile(config.SlotsFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var slots []entities.Slot

	// Decode existing slots from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&slots); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing results:", err)
		return err
	}

	for _, slot := range slots {
		if slot.Date == date {
			for _, game := range slot.Games {
				if game.GameID == gameId {
					for _, v := range game.Slots {
						if v.Time == time {
							v.BookedBy = append(v.BookedBy, user)
						}
					}
				}
			}
		}
	}

	// Truncate the file to overwrite it with the updated slots array
	if err := file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}

	// Move the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return err
	}

	// Encode the updated slots array to JSON and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(slots); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	return nil
}

func (r *slotRepo) GetPendingInvites(user entities.User, date string) ([]entities.Invites, error) {
	var invites []entities.Invites

	// Traverse through all games for the given date
	for gameID, slots := range globals.SlotsMap[date] {
		// Traverse through all slots for the current game
		for _, slot := range slots {
			// Check if the user is in the InvitedUsers array
			for _, invitedUser := range slot.InvitedUsers {
				if invitedUser.UserId == user.UserId {
					// User is invited; add the invite details to the invites slice
					invites = append(invites, entities.Invites{
						Date: date,
						Game: globals.GamesMap[gameID].Name,
						Time: slot.Time,
					})
					break
				}
			}
		}
	}

	return invites, nil
}

func loadAllSlots() {
	return
}
