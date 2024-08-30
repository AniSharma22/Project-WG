package service_test

import (
	"project2/pkg/validation"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"john.doe@watchguard.com", true},
		{"Jane.Doe@watchguard.com", true},
		{"jane.doe@gmail.com", false},
		{"john@watchguard.com", false},
		{"john.doe@watchguard.co", false},
	}

	for _, test := range tests {
		if result := validation.IsValidEmail(test.email); result != test.expected {
			t.Errorf("IsValidEmail(%s) = %v; want %v", test.email, result, test.expected)
		}
	}
}

func TestIsValidPhoneNumber(t *testing.T) {
	tests := []struct {
		phoneNumber string
		expected    bool
	}{
		{"9876543210", true},
		{"6789012345", true},
		{"5123456789", false},  // Does not start with 6, 7, 8, or 9
		{"987654321", false},   // Less than 10 digits
		{"98765432101", false}, // More than 10 digits
	}

	for _, test := range tests {
		if result := validation.IsValidPhoneNumber(test.phoneNumber); result != test.expected {
			t.Errorf("IsValidPhoneNumber(%s) = %v; want %v", test.phoneNumber, result, test.expected)
		}
	}
}

func TestIsValidGender(t *testing.T) {
	tests := []struct {
		gender   string
		expected bool
	}{
		{"male", true},
		{"female", true},
		{"other", true},
		{"Male", true},
		{"males", false}, // Invalid gender
		{"", false},      // Empty string
	}

	for _, test := range tests {
		if result := validation.IsValidGender(test.gender); result != test.expected {
			t.Errorf("IsValidGender(%s) = %v; want %v", test.gender, result, test.expected)
		}
	}
}
