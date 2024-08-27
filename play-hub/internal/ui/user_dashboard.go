package ui

import (
	"fmt"
	"strings"
)

func (ui *UI) ShowUserDashboard() {
	for {
		fmt.Println("\033[1;33m") // Yellow bold (close to orange)
		fmt.Println("╔═════════════════════════════════════╗")
		fmt.Println("║          🎮 User Dashboard 🎮       ║")
		fmt.Println("╚═════════════════════════════════════╝")
		fmt.Println("\033[0m") // Reset color

		fmt.Println("Please choose an option:")
		fmt.Println("1. Game Room")
		fmt.Println("2. View Pending Invites")
		fmt.Println("3. View Leaderboard")
		fmt.Println("4. Update Results")
		fmt.Println("5. View Upcoming Bookings")
		fmt.Println("6. View Profile")
		fmt.Println("7. Logout")

		fmt.Print("Enter your choice (1-7): ")
		choice, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ui.ShowGameRoom() // Placeholder function, implement as needed
		case "2":
			ui.ViewPendingInvites() // Placeholder function, implement as needed
		case "3":
			ui.ViewLeaderboard() // Placeholder function, implement as needed
		case "4":
			ui.UpdateResults()
		case "5":
			ui.ViewUpcomingBookings()
		case "6":
			ui.ViewProfile()
		case "7":
			fmt.Println("Logging out...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		}
	}
}
