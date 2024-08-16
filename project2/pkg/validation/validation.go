package validation

import (
	"project2/pkg/globals"
	"regexp"
	"strings"
	"unicode"
)

func IsValidUsername(username string) bool {
	if len(username) == 0 || len(username) > 30 {
		return false
	}
	return true
}

// checks if the email is of the format name.surname@watchguard.com
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+\.[a-zA-Z]+@watchguard\.com$`)
	return re.MatchString(strings.ToLower(email))
}

func EmailAlreadyExists(email string) bool {
	_, exists := globals.UsersMap[email]
	return exists
}

// Password must have 1 uppercase, 1 lowercase, 1 special character and minimum 8 length
func IsValidPassword(password string) bool {
	var hasUpper, hasLower, hasSpecial bool

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasSpecial
}

// checks if length of number is 10 and starts with 6,7,8 or 9
func IsValidPhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^[6-9]\d{9}$`)
	return re.MatchString(phoneNumber)
}

func IsAdmin(email string) bool {
	return globals.UsersMap[email].Role == "admin"
}

func IsValidGender(gender string) bool {
	gender = strings.ToLower(gender)
	return gender == "male" || gender == "female" || gender == "other"
}
