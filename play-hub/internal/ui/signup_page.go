package ui

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"project2/internal/domain/entities"
	"project2/pkg/utils"
	"project2/pkg/validation"
	"strings"
	"syscall"
)

func (ui *UI) ShowSignupPage() {
	var email, password, phoneNo, gender string

	// Get valid email
	for {
		fmt.Print("Enter your email: ")
		email, _ = ui.reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if validation.IsValidEmail(email) && ui.userService.EmailAlreadyExists(email) {
			break
		} else {
			fmt.Println("Invalid email. Please try again.")
		}
	}

	// Get and confirm password
	for {
		fmt.Println("(1 Capital, 1 small, 1 special character with min 8 length)")
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
		} else if !validation.IsValidPassword(string(bytePassword1)) {
			fmt.Println("Password complexity not met!")
		} else {
			password, _ = utils.GetHashedPassword(bytePassword1)
			break
		}
	}

	// Get valid phone number
	for {
		fmt.Print("Enter your phone number: ")
		phoneNo, _ = ui.reader.ReadString('\n')
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
		gender, _ = ui.reader.ReadString('\n')
		gender = strings.TrimSpace(gender)
		if validation.IsValidGender(gender) {
			break
		} else {
			fmt.Println("Invalid gender. Please try again.")
		}
	}

	// Create the user entity
	user := entities.User{
		Email:    email,
		Password: password,
		PhoneNo:  phoneNo,
		Gender:   gender,
	}

	// Sign up the user
	if err := ui.userService.Signup(&user); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Signup successful!")
	ui.ShowUserDashboard()
}
