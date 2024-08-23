package utils

import (
	"fmt"
	"project/internal/models"
	"reflect"
	"strings"
)

var (
	TodosMap         = make(map[string]models.UserDetails)
	ProgressMap      = make(map[string][]models.CompletedSection)
	CourseOutlineMap = make(map[int]models.CourseDetails)
	UserMap          = make(map[string]string)
	UserDataLoaded   bool
)

func IsValidCountry(country string) bool {
	country = strings.ToLower(country)
	if country == "pakistan" || country == "china" {
		return false
	}

	return true
}

// IsUsernameTaken checks if the username is taken, ensuring that data has been loaded.
func IsUsernameTaken(username string) bool {

	if !UserDataLoaded {
		// Data not loaded, return false as we can't confirm the existence of the username
		return false
	}

	_, userExists := UserMap[username]
	return userExists
}

func LoadUsers(filename string) {
	userDataChan := make(chan any)
	go RfileGeneral(userDataChan, filename, reflect.TypeOf(models.UserData{}))

	for user := range userDataChan {
		// Type assertion
		userData, ok := user.(*models.UserData)
		if !ok {
			fmt.Println("Error: received data is not of type models.UserData")
			continue
		}

		UserMap[userData.Username] = userData.Password
	}
	UserDataLoaded = true
}

func EnsureDataLoaded() error {

	if !UserDataLoaded {
		return fmt.Errorf("data not yet loaded")
	}
	return nil
}

func LoadCourseOutline(filename string) {
	courseDataChan := make(chan any)
	go RfileGeneral(courseDataChan, filename, reflect.TypeOf(models.CourseOutline{}))

	for course := range courseDataChan {
		courseData, ok := course.(*models.CourseOutline)
		if !ok {
			fmt.Println("Error: received data is not of type models.Course")
		}
		CourseOutlineMap[courseData.ID] = courseData.Details

	}

}

func LoadUserProgress(filename string) {
	progressChan := make(chan any)

	go RfileGeneral(progressChan, filename, reflect.TypeOf(models.UserProgress{}))

	for progress := range progressChan {
		userProgress, ok := progress.(*models.UserProgress)
		if !ok {
			fmt.Println("Error: received data is not of type models.Course")
		}

		ProgressMap[userProgress.Username] = userProgress.Progress

	}

}
