package ui

import (
	"fmt"
	"project2/pkg/globals"
)

func (ui *UI) ViewProfile() {
	fmt.Println("profile")
	user, err := ui.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return
	}
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Gender: %v\n", user.Gender)
	fmt.Printf("Phone Number: %v\n", user.PhoneNo)
	fmt.Printf("Role: %s\n", user.Role)
	fmt.Printf("Games Played: %d\n", user.Wins+user.Losses)
	fmt.Printf("Score: %d\n", user.OverallScore)
}
