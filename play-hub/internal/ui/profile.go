package ui

import (
	"context"
	"fmt"
	"project2/pkg/globals"
)

func (ui *UI) ViewProfile() {
	fmt.Println("👤  Your Profile  👤")

	user, err := ui.userService.GetUserByID(context.Background(), globals.ActiveUser)
	if err != nil {
		fmt.Println("⚠️ Error fetching profile:", err)
		return
	}

	fmt.Println("------------------------------------------------")
	fmt.Printf("📧  Name:        %s\n", user.Username)
	fmt.Printf("📧  Email:        %s\n", user.Email)
	fmt.Printf("🚻  Gender:       %v\n", user.Gender)
	fmt.Printf("📱  Phone Number: %v\n", user.MobileNumber)
	fmt.Printf("🎭  Role:         %s\n", user.Role)
	fmt.Println("------------------------------------------------")
}
