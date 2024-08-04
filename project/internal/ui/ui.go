package ui

import (
	"fmt"
	"project/internal/auth"
	"project/internal/utils"
)

// HandleUserAction processes user actions for signup and login.
func HandleUserAction(method string) {
	var username, password, country string

	// Read the username
	fmt.Print("Enter username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	if method == "signup" {
		// Ensure data is loaded before proceeding
		for {
			err = utils.EnsureDataLoaded()
			if err == nil {
				break
			}
		}
		// Check if the username is already taken
		for utils.IsUsernameTaken(username) {
			fmt.Println("Username already taken!!")
			fmt.Print("Enter a new Username: ")
			_, err = fmt.Scanln(&username)
			if err != nil {
				fmt.Println("Error reading username:", err)
				return
			}
		}
	}

	// Read the password
	fmt.Println("(1 Capital, 1 small, 1 special charcter with min 8 length)")
	fmt.Print("Enter password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	fmt.Println("Password entered.") // Debug print

	if method == "signup" && !utils.IsValidPassword(password) {
		// Prompt for a valid password if the provided one is not strong enough
		for {
			fmt.Println("Enter a stronger password!")
			var tempPass string
			_, err = fmt.Scanln(&tempPass)
			if err != nil {
				fmt.Println("Error reading password:", err)
				continue
			}
			if utils.IsValidPassword(tempPass) {
				password = tempPass
				break
			} else {
				fmt.Println("Password is still not strong enough. Try again.")
			}
		}
	}

	if method == "signup" {
		// Read the country
		fmt.Print("Enter country: ")
		_, err = fmt.Scanln(&country)
		if err != nil {
			fmt.Println("Error reading country:", err)
			return
		}
		fmt.Println("Country entered:", country) // Debug print
		if !utils.IsValidCountry(country) {
			fmt.Println("Users from this country are not allowed!!")
			return
		}
	}

	// Handle the action based on the method
	switch method {
	case "login":
		auth.Login(username, password)
	case "signup":
		auth.Signup(username, password)
	default:
		fmt.Println("Invalid method. Please use 'login' or 'signup'.")
	}
}
