package main

import (
	"fmt"
	"os"
	"project/internal/models"
	"project/internal/ui"
	"project/internal/utils"
	"reflect"
)

func init() {
	// Check if the users.json file exists and load users if it does
	_, err := os.Stat("users.json")
	if err == nil {
		go utils.LoadUsers()
	} else {
		utils.DataLoaded = true
	}
	if _, err := os.Stat("progress.json"); err == nil {
		go utils.LoadUserProgress()
	}
	// Start loading the course outline
	go utils.LoadCourseOutline()

}

func main() {
	for {
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
			fmt.Println("------------")
			fmt.Println("Signup")
			fmt.Println("------------")
			ui.HandleUserAction("signup")
		case "2":
			fmt.Println("------------")
			fmt.Println("Login")
			fmt.Println("------------")
			ui.HandleUserAction("login")
		case "3":
			if utils.NewEntryAdded {
				fmt.Println("Saving Data to File ...")
				utils.WfileGeneral("users.json", utils.UserMap, reflect.TypeOf(models.UserData{}))
			}
			if utils.ProgressMap != nil {
				utils.WfileGeneral("progress.json", utils.ProgressMap, reflect.TypeOf(models.UserProgress{}))
			}

			fmt.Println("Exiting...")

			return
		default:
			fmt.Println("Invalid choice. Please enter 1 for Signup, 2 for Login, or 3 for Exit.")
		}
	}
}
