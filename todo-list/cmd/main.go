package main

import (
	"fmt"
	"os"
	"project/internal/config"
	"project/internal/models"
	"project/internal/ui"
	"project/internal/utils"
	"reflect"
)

func init() {
	// Check if the users.json file exists and load users if it does
	_, err := os.Stat(config.UsersFile)
	if err == nil {
		go utils.LoadUsers(config.UsersFile)
	} else {
		utils.UserDataLoaded = true
	}

	if _, err := os.Stat(config.ProgressFile); err == nil {
		go utils.LoadUserProgress(config.ProgressFile)
	}

	if _, err := os.Stat(config.TodosFile); err == nil {
		go utils.LoadUserTodos(config.TodosFile)
	}

	// Start loading the course outline
	go utils.LoadCourseOutline(config.CoursesFile)

}

func main() {
	for {
		var choice string
		fmt.Println("\033[1;36m") // Cyan bold
		fmt.Println("===================================")
		fmt.Println("     		  WELCOME    ")
		fmt.Println("===================================")
		fmt.Println("\033[0m") // Reset color
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
			if utils.ProgressMap != nil {
				utils.WfileGeneral(config.ProgressFile, utils.ProgressMap, reflect.TypeOf(models.UserProgress{}))
			}
			if utils.TodosMap != nil {
				utils.WfileGeneral(config.TodosFile, utils.TodosMap, reflect.TypeOf(models.UserTodos{}))
			}
			fmt.Println("Exiting...")

			return
		default:
			fmt.Println("Invalid choice. Please enter 1 for Signup, 2 for Login, or 3 for Exit.")
		}
	}
}
