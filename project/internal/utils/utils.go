package utils

import (
	"fmt"
	"project/internal/models"
	"strings"
	"sync"
	"unicode"
)

var (
	UserMap       = make(map[string]string)
	userMapMutex  sync.RWMutex
	dataLoaded    bool
	dataLoadedMu  sync.Mutex
	NewEntryAdded bool = false
)

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

func IsValidCountry(country string) bool {
	country = strings.ToLower(country)
	if country == "pakistan" || country == "china" {
		fmt.Println("Users from this country are not allowed!!")
		return false
	}

	return true
}

func LoadUsers() {
	userDataChan := make(chan models.UserData)
	go Rfile(userDataChan)

	for user := range userDataChan {
		userMapMutex.Lock()
		UserMap[user.Username] = user.Password
		userMapMutex.Unlock()
	}

	// Mark data as loaded
	userMapMutex.Lock()
	dataLoaded = true
	userMapMutex.Unlock()
}

func EnsureDataLoaded() error {
	dataLoadedMu.Lock()
	defer dataLoadedMu.Unlock()

	if !dataLoaded {
		return fmt.Errorf("data not yet loaded")
	}
	return nil
}

// IsUsernameTaken checks if the username is taken, ensuring that data has been loaded.
func IsUsernameTaken(username string) bool {
	userMapMutex.RLock()
	defer userMapMutex.RUnlock()

	if !dataLoaded {
		// Data not loaded, return false as we can't confirm the existence of the username
		return false
	}

	_, userExists := UserMap[username]
	return userExists
}
