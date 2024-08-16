package ui

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"project2/internal/domain/entities"
	"project2/pkg/validation"
	"strings"
	"syscall"
)

func (ui *UI) ShowSignupPage() {
	reader := bufio.NewReader(os.Stdin)
	var username, email, password, phoneNo, gender string

	// Get valid username
	for {
		fmt.Print("Enter your username: ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if validation.IsValidUsername(username) {
			break
		} else {
			fmt.Println("Invalid username. Please try again.")
		}
	}

	// Get valid email
	for {
		fmt.Print("Enter your email: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if validation.IsValidEmail(email) && !validation.EmailAlreadyExists(email) {
			break
		} else {
			fmt.Println("Invalid email. Please try again.")
		}
	}

	for {
		fmt.Print("Enter your password: ")
		bytePassword1, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading password:", err)
		}
		fmt.Println() // Add a newline after password input

		fmt.Print("Confirm your password: ")
		bytePassword2, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading password:", err)
		}
		fmt.Println() // Add a newline after password input

		if string(bytePassword1) != string(bytePassword2) {
			fmt.Println("Passwords did not match. Please try again.")
		} else {
			password = string(bytePassword1)
			break
		}
	}

	// Get valid phone number
	for {
		fmt.Print("Enter your phone number: ")
		phoneNo, _ = reader.ReadString('\n')
		phoneNo = strings.TrimSpace(phoneNo)
		if validation.IsValidPhoneNumber(phoneNo) {
			break
		} else {
			fmt.Println("Invalid phone number. Please try again.")
		}
	}

	// Get valid gender
	for {
		fmt.Print("Enter your gender (Male/Female/Other): ")
		gender, _ = reader.ReadString('\n')
		gender = strings.TrimSpace(gender)
		if validation.IsValidGender(gender) {
			break
		} else {
			fmt.Println("Invalid gender. Please try again.")
		}
	}

	user := entities.User{
		Name:     username,
		Email:    email,
		Password: password,
		PhoneNo:  phoneNo,
		Gender:   gender,
	}

	if err := ui.userService.Signup(&user); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Signup successful!")
	ui.ShowUserDashboard()
}
