package ui

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"project2/pkg/validation"
	"strings"
	"syscall"
)

func (ui *UI) ShowLoginPage() {
	reader := bufio.NewReader(os.Stdin)

	var email string
	var password string

	// Get valid email
	for {
		fmt.Print("Enter your email: ")
		email, _ = reader.ReadString('\n')
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
	password = string(bytePassword)
	fmt.Println() // Add a newline after password input

	// Verify email and password
	user, exists := globals.UsersMap[email]
	if !exists {
		fmt.Println("User does not exist.")
		return
	}

	if !utils.VerifyPassword(password, user.Password) {
		fmt.Println("Invalid password.")
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
