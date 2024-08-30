package service_test

import (
	"project2/pkg/utils"
	"project2/pkg/validation"
	"testing"
)

func TestGetHashedPassword(t *testing.T) {
	password := "Password123!"

	hash, err := utils.GetHashedPassword([]byte(password))
	if err != nil {
		t.Errorf("GetHashedPassword returned an error: %v", err)
	}

	if len(hash) == 0 {
		t.Error("Expected non-empty hash, got empty string")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "Password123!"
	wrongPassword := "WrongPassword!"

	hash, err := utils.GetHashedPassword([]byte(password))
	if err != nil {
		t.Fatalf("GetHashedPassword returned an error: %v", err)
	}

	if !utils.VerifyPassword([]byte(password), hash) {
		t.Error("VerifyPassword returned false for correct password")
	}

	if utils.VerifyPassword([]byte(wrongPassword), hash) {
		t.Error("VerifyPassword returned true for incorrect password")
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Password1!", true},
		{"Pass!", false},      // Less than 8 characters
		{"password1!", false}, // No uppercase letter
		{"PASSWORD1!", false}, // No lowercase letter
		{"Password1", false},  // No special character
		{"Password!@", true},  // Two special characters
	}

	for _, test := range tests {
		if result := validation.IsValidPassword(test.password); result != test.expected {
			t.Errorf("IsValidPassword(%s) = %v; want %v", test.password, result, test.expected)
		}
	}
}

//func TestGetTotalScore(t *testing.T) {
//	tests := []struct {
//		totalWins   int
//		totalLosses int
//		expected    float32
//	}{
//		{10, 0, 1.0},
//		{10, 10, 0.11},
//		{20, 5, 0.56},
//		{0, 0, 0},
//	}
//
//	for _, test := range tests {
//		result := utils.GetTotalScore(test.totalWins, test.totalLosses)
//		if result != test.expected {
//			t.Errorf("GetTotalScore(%d, %d) = %v; want %v", test.totalWins, test.totalLosses, result, test.expected)
//		}
//	}
//}

func TestGetNameFromEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected string
	}{
		{"john.doe@example.com", "john doe"},
		{"jane.smith@domain.com", "jane smith"},
		{"user@domain.com", "user"},
	}

	for _, test := range tests {
		result := utils.GetNameFromEmail(test.email)
		if result != test.expected {
			t.Errorf("GetNameFromEmail(%s) = %v; want %v", test.email, result, test.expected)
		}
	}
}
