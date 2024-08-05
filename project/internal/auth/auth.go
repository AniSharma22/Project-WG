package auth

import (
	"fmt"
	"project/internal/utils"
	"strconv"
)

//const userFile = "users.json"

func Signup(username, password string) {
	utils.UserMap[username] = password
	utils.NewEntryAdded = true
	utils.ProgressMap[username] = []int{}
	fmt.Println("Signup successful!")
	HandleUserDashboard(username)
}

func Login(username, password string) {
	if utils.IsUsernameTaken(username) && utils.UserMap[username] == password {
		fmt.Println("Login successful!")
		HandleUserDashboard(username)
		return
	}
	fmt.Println("Username or Password is wrong!!")
	return
}

func HandleUserDashboard(username string) {
	for {
		// Ensure CourseOutline is loaded
		if len(utils.CourseOutline) == 0 {
			fmt.Println("Error: CourseOutline is empty or not loaded.")
			return
		}

		mp := make(map[int]bool, 20)
		var userChoice string

		// Fill mp with completed modules
		for _, val := range utils.ProgressMap[username] {
			mp[val] = true
		}

		// Display user progress
		fmt.Println(username, "                  ", len(utils.ProgressMap[username]), "/10 Modules Done")
		fmt.Println(len(utils.ProgressMap[username])*10, "% Completed")
		fmt.Println()
		fmt.Println("--------------------------------------------------------")

		// Find and display uncompleted modules
		for i := 1; i <= 10; i++ {
			if mp[i] {
				fmt.Println(i, utils.CourseOutline[1].Modules[i-1].Title, "✅")
			} else {
				fmt.Println(i, utils.CourseOutline[1].Modules[i-1].Title, "❌")
			}
		}

		fmt.Println("--------------------------------------------------------")
		fmt.Println("Press 0 to Logout or enter the module ID to mark it as completed")
		fmt.Print("Enter: ")
		_, _ = fmt.Scanln(&userChoice)

		switch userChoice {
		case "0":
			fmt.Println("Logout successful!")
			return
		default:
			// Attempt to mark module as completed
			moduleID, err := strconv.Atoi(userChoice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid module ID or '0' to logout.")
				continue
			}

			if moduleID < 1 || moduleID > 10 {
				fmt.Println("Invalid module ID. Please enter a valid module ID.")
			} else if mp[moduleID] {
				fmt.Println("Module already completed.")
			} else {
				// Add moduleID to completedModules
				utils.ProgressMap[username] = append(utils.ProgressMap[username], moduleID)
				fmt.Println("Module marked as completed!")
			}
		}
	}
}
