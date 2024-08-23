package ui

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"project2/pkg/validation"
	"strings"
	"syscall"
)

func (ui *UI) ShowLoginPage() {

	var email string
	var password []byte

	// Get valid email
	for {
		fmt.Print("Enter your email: ")
		email, _ = ui.reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if validation.IsValidEmail(email) {
			break
		} else {
			fmt.Println("Invalid email. Please try again.")
		}
	}

	// Get password securely
	fmt.Print("Enter your password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	password = bytePassword
	fmt.Println() // Add a newline after password input

	user, err := ui.userService.Login(email, password)
	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}

	fmt.Println("Login successful.")
	// Redirect to appropriate dashboard
	if user.Role == "admin" {
		ui.ShowAdminDashboard()
	} else {
		ui.ShowUserDashboard()
	}
}
