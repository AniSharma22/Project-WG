package ui

import (
	"fmt"
	"strings"
)

func (ui *UI) ShowMainMenu() {
	for {
		fmt.Println("\033[1;36m") // Cyan bold
		fmt.Println("===================================")
		fmt.Println("     		  WELCOME    ")
		fmt.Println("===================================")
		fmt.Println("\033[0m") // Reset color
		fmt.Println("Please choose an option:")
		fmt.Println("1. Signup")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice (1, 2, or 3): ")

		// Read the input using the bufio.Reader
		choice, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim any whitespace or newline characters
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("------------")
			fmt.Println("Signup")
			fmt.Println("------------")
			ui.ShowSignupPage()
		case "2":
			fmt.Println("------------")
			fmt.Println("Login")
			fmt.Println("------------")
			ui.ShowLoginPage()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1 for Signup, 2 for Login, or 3 for Exit.")
		}
	}
}
