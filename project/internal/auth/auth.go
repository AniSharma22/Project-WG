package auth

import (
	"fmt"
	"project/internal/utils"
)

const userFile = "users.json"

func Signup(username, password string) {
	utils.UserMap[username] = password
	utils.NewEntryAdded = true
	fmt.Println("Signup successful!")
}

func Login(username, password string) {
	if utils.IsUsernameTaken(username) && utils.UserMap[username] == password {
		fmt.Println("Login successful!")
		return
	}
	fmt.Println("Username or Password is wrong!!")
	return
}
