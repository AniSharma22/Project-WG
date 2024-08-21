package ui

import (
	"fmt"
	"strings"
)

func (ui *UI) ShowUserDashboard() {
	for {
		fmt.Println("===================================")
		fmt.Println("          User Dashboard           ")
		fmt.Println("===================================")
		fmt.Println("Please choose an option:")
		fmt.Println("1. Game Room")
		fmt.Println("2. View Pending Invites")
		fmt.Println("3. View Leaderboard")
		fmt.Println("4. View Profile")
		fmt.Println("5. Logout")

		fmt.Print("Enter your choice (1-5): ")
		choice, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Redirecting to Game Room...")
			ui.ShowGameRoom() // Placeholder function, implement as needed
		case "2":
			fmt.Println("Viewing Pending Invites...")
			ui.ViewPendingInvites() // Placeholder function, implement as needed
		case "3":
			fmt.Println("Viewing Leaderboard...")
			ui.ViewLeaderboard() // Placeholder function, implement as needed
		case "4":
			fmt.Println("Viewing Profile...")
			ui.ViewProfile() // Placeholder function, implement as needed
		case "5":
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
