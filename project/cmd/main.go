package main

import (
	"fmt"
	"os"
	"project/internal/ui"
	"project/internal/utils"
)

func init() {
	if _, err := os.Stat("users.json"); err == nil {
		go utils.LoadUsers()
	}
}

func main() {
	for 
		var choice string

		fmt.Println("Please choose an option:")
		fmt.Println("1. Signup")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice (1, 2 or 3): ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case "1":
			fmt.Println("Signup")
			ui.HandleUserAction("signup")
		case "2":
			fmt.Println("Login")
			ui.HandleUserAction("login")
		case "3":
			if utils.NewEntryAdded {
				fmt.Println("Saving Data to File ...")
				utils.Wfile(utils.UserMap)
			}
			fmt.Println("Exiting...")

			return
		default:
			fmt.Println("Invalid choice. Please enter 1 for Signup, 2 for Login, or 3 for Exit.")
		}
	}
}
