package ui

import (
	"fmt"
	"project/internal/ui/course"
	"project/internal/ui/dailyStatus"
	"project/internal/ui/profile"
	"project/internal/ui/todo"
)

// HandleUserDashboard displays the dashboard and processes user options
func HandleUserDashboard(username string) {
	for {
		fmt.Println("\033[1;32m") // Cyan bold
		fmt.Println("===============================")
		fmt.Println("          DASHBOARD    ")
		fmt.Println("===============================")
		fmt.Println("1. Manage courses")
		fmt.Println("2. Manage todos")
		fmt.Println("3. View daily status")
		fmt.Println("4. View profile")
		fmt.Println("5. Logout")
		fmt.Println("===============================")
		fmt.Println("\033[0m") // Reset color
		fmt.Print("Choose an option: ")

		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case "1":
			// Implement course management logic
			course.ManageCourses(username)
		case "2":
			// Implement t.odo management logic
			todo.ManageTodos(username)
		case "3":
			// Implement view daily status logic
			dailyStatus.ViewDailyStatus(username)
		case "4":
			// Implement view profile logic
			profile.ViewProfile(username)
		case "5":
			fmt.Println("Logging out...")
			return // Exit the dashboard loop and log out
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}
