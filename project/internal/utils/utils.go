package utils

import (
	"fmt"
	"project/internal/models"
	"reflect"
	"strings"
	"unicode"
)

var (
	ProgressMap   = make(map[string][]int)
	CourseOutline = make([]models.Course, 1)
	UserMap       = make(map[string]string)
	DataLoaded    bool
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
	userDataChan := make(chan any)
	go RfileGeneral(userDataChan, "users.json", reflect.TypeOf(models.UserData{}))

	for user := range userDataChan {
		// Type assertion
		userData, ok := user.(*models.UserData)
		if !ok {
			fmt.Println("Error: received data is not of type models.UserData")
			continue
		}

		UserMap[userData.Username] = userData.Password
	}
	DataLoaded = true
}

func EnsureDataLoaded() error {

	if !DataLoaded {
		return fmt.Errorf("data not yet loaded")
	}
	return nil
}

// IsUsernameTaken checks if the username is taken, ensuring that data has been loaded.
func IsUsernameTaken(username string) bool {

	if !DataLoaded {
		// Data not loaded, return false as we can't confirm the existence of the username
		return false
	}

	_, userExists := UserMap[username]
	return userExists
}

func LoadCourseOutline() {
	courseDataChan := make(chan any)
	go RfileGeneral(courseDataChan, "courses.json", reflect.TypeOf(models.Course{}))

	for course := range courseDataChan {
		courseData, ok := course.(*models.Course)
		if !ok {
			fmt.Println("Error: received data is not of type models.Course")
		}
		CourseOutline = append(CourseOutline, *courseData)

	}

}

func LoadUserProgress() {
	progressChan := make(chan any)

	go RfileGeneral(progressChan, "progress.json", reflect.TypeOf(models.UserProgress{}))

	for progress := range progressChan {
		userProgress, ok := progress.(*models.UserProgress)
		if !ok {
			fmt.Println("Error: received data is not of type models.Course")
		}
		ProgressMap[userProgress.Username] = userProgress.CompletedModules

	}

}
