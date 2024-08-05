package utils

import "testing"

func TestIsValidPassword(t *testing.T) {
	invalidPasswords := []string{"anish", "Anish"}
	validPasswords := []string{"Random20394@!", "%Shwoirwer2"}

	for _, password := range invalidPasswords {
		if IsValidPassword(password) {
			t.Errorf("Password '%s' should be invalid", password)
		}
	}

	for _, password := range validPasswords {
		if !IsValidPassword(password) {
			t.Errorf("Password '%s' should be valid", password)
		}
	}
}

func TestIsValidCountry(t *testing.T) {
	validCountries := []string{"usa", "canada", "india", "nepal"}
	invalidCountries := []string{"pakistan", "china"}

	for _, country := range validCountries {
		if !IsValidCountry(country) {
			t.Errorf("Country '%s' should be valid", country)
		}
	}

	for _, country := range invalidCountries {
		if IsValidCountry(country) {
			t.Errorf("Country '%s' should be invalid", country)
		}
	}
}
