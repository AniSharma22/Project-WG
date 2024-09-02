package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"project2/pkg/globals"
	"strconv"
	"strings"
)

func (ui *UI) UpdateResults() {
	fmt.Println("\n=============================== Results to Update ===============================")

	gameHistoryList, err := ui.bookingService.GetBookingsToUpdateResult(context.Background(), globals.ActiveUser)
	if err != nil {
		fmt.Printf("Error retrieving results: %v\n", err)
		return
	}

	if len(gameHistoryList) == 0 {
		fmt.Println("No results to update.")
		return
	}

	for i, gameHistory := range gameHistoryList {
		//game, _ := ui.gameService.GetGameByID(gameHistory.GameID)
		//slot, _ := ui.slotService.GetSlotById(gameHistory.SlotID)

		fmt.Printf(" #%d\n", i+1)
		fmt.Printf("Game:         %s\n", gameHistory.GameName)
		fmt.Printf("Start Time:   %s IST\n", gameHistory.StartTime.Format("03:04 PM"))
		fmt.Printf("End Time:     %s IST\n", gameHistory.EndTime.Format("03:04 PM"))

		if len(gameHistory.BookedUsers) > 0 {
			fmt.Println("Participants:")
			for _, userName := range gameHistory.BookedUsers {
				fmt.Printf("- %s\n", userName)
			}
		} else {
			fmt.Println("Participants: None")
		}

		fmt.Println(strings.Repeat("-", 80))
	}

	fmt.Println("Press the corresponding number to update the result of that game:")

	// Read user input for selection
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Parse the input to get the index
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(gameHistoryList) {
		fmt.Println("Invalid selection. Please enter a valid number.")
		return
	}

	// Adjust index to match slice (0-based indexing)
	index = index - 1

	// Confirm the selection and prompt for result update
	selectedGameHistory := gameHistoryList[index]
	fmt.Printf("Selected Game: %s\n", selectedGameHistory.BookingId)
	fmt.Println("Press 'w' for Win or 'l' for Loss:")

	resultInput, _ := reader.ReadString('\n')
	resultInput = strings.TrimSpace(strings.ToUpper(resultInput))

	switch resultInput {
	case "W":
		// Update the result as a win
		err := ui.bookingService.UpdateBookingResult(context.Background(), selectedGameHistory.BookingId, "win")
		if err != nil {
			fmt.Printf("Error updating result: %v\n", err)
		} else {
			fmt.Println("Result updated to Win!")
		}
	case "L":
		// Update the result as a loss
		err := ui.bookingService.UpdateBookingResult(context.Background(), selectedGameHistory.BookingId, "loss")
		if err != nil {
			fmt.Printf("Error updating result: %v\n", err)
		} else {
			fmt.Println("Result updated to Loss!")
		}
	default:
		fmt.Println("Invalid input. Please enter 'W' or 'L'.")
	}
}
