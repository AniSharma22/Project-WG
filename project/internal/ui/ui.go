package ui

import (
	"bufio"
	"fmt"
	"os"
	"project/internal/auth"
	"project/internal/utils"
	"strings"
)

// HandleUserAction processes user actions for signup and login.
func HandleUserAction(method string) {
	reader := bufio.NewReader(os.Stdin)
	var username, password, country string

	// Read the username
	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username) // Trim any extra spaces or newlines

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
			username, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading username:", err)
				return
			}
			username = strings.TrimSpace(username) // Trim any extra spaces or newlines
		}
	}

	// Read the password
	fmt.Println("(1 Capital, 1 small, 1 special character with min 8 length)")
	fmt.Print("Enter password: ")
	password, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	password = strings.TrimSpace(password) // Trim any extra spaces or newlines

	if method == "signup" && !utils.IsValidPassword(password) {
		// Prompt for a valid password if the provided one is not strong enough
		for {
			fmt.Println("Enter a stronger password!")
			tempPass, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading password:", err)
				continue
			}
			tempPass = strings.TrimSpace(tempPass) // Trim any extra spaces or newlines
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
		country, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading country:", err)
			return
		}
		country = strings.TrimSpace(country) // Trim any extra spaces or newlines
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
