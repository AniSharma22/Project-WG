package ui

import (
	"fmt"
	"project2/pkg/utils"
	"strings"
)

func (ui *UI) ViewUpcomingBookings() {
	fmt.Println("\n=============================== Your Upcoming Bookings ===============================\n")

	slots, err := ui.slotService.GetUpcomingBookedSlots()
	if err != nil {
		fmt.Printf("Error retrieving bookings: %v\n", err)
		return
	}

	if len(slots) == 0 {
		fmt.Println("You have no upcoming bookings.")
		return
	}

	for i, slot := range slots {
		game, _ := ui.gameService.GetGameByID(slot.GameID)

		fmt.Printf("Booking #%d\n", i+1)
		fmt.Printf("Game:         %s\n", game.Name)
		fmt.Printf("Start Time:   %s IST\n", slot.StartTime.Format("03:04 PM"))
		fmt.Printf("End Time:     %s IST\n", slot.EndTime.Format("03:04 PM"))

		if len(slot.BookedUsers) > 0 {
			fmt.Println("Participants: ")
			for _, userID := range slot.BookedUsers {
				user, _ := ui.userService.GetUserById(userID)
				fmt.Printf("- %s\n", utils.GetNameFromEmail(user.Email))
			}
		} else {
			fmt.Println("Participants: None")
		}

		if i < len(slots)-1 {
			fmt.Println(strings.Repeat("-", 80))
		}
	}

	fmt.Println("\n======================================================================================\n")
}
