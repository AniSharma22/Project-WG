package ui

import (
	"context"
	"fmt"
	"project2/pkg/globals"
	"strings"
)

func (ui *UI) ViewUpcomingBookings() {
	fmt.Println("\n=============================== Your Upcoming Bookings ===============================")

	bookings, err := ui.bookingService.GetUpcomingBookings(context.Background(), globals.ActiveUser)
	if err != nil {
		fmt.Printf("Error retrieving bookings: %v\n", err)
		return
	}

	if len(bookings) == 0 {
		fmt.Println("You have no upcoming bookings.")
		return
	}

	for i, booking := range bookings {
		//game, _ := ui.gameService.GetGameByID(slot.GameID)

		fmt.Printf("Booking #%d\n", i+1)
		fmt.Printf("Game:         %s\n", booking.GameName)
		fmt.Printf("Start Time:   %s IST\n", booking.StartTime.Format("03:04 PM"))
		fmt.Printf("End Time:     %s IST\n", booking.EndTime.Format("03:04 PM"))

		if len(booking.BookedUsers) > 0 {
			fmt.Println("Participants: ")
			for _, name := range booking.BookedUsers {
				fmt.Printf("- %s\n", name)
			}
		} else {
			fmt.Println("Participants: None")
		}

		if i < len(bookings)-1 {
			fmt.Println(strings.Repeat("-", 80))
		}
	}

	fmt.Println("\n======================================================================================")
}
