package utils

import (
	"project/internal/models"
	"reflect"
	"testing"
)

const testUsersFile = "C:\\Users\\anisharma\\GolandProjects\\goprac\\project\\testJson\\testUsers.json"
const testCoursesFile = "C:\\Users\\anisharma\\GolandProjects\\goprac\\project\\testJson\\testCourses.json"
const testProgressFile = "C:\\Users\\anisharma\\GolandProjects\\goprac\\project\\testJson\\testProgress.json"
const testTodosFile = "C:\\Users\\anisharma\\GolandProjects\\goprac\\project\\testJson\\testTodos.json"

// Mock data for testing
var (
	mockUserMap = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
	mockProgressMap = map[string][]models.CompletedSection{
		"user1": {
			{CourseID: 1, CompletedSections: []float64{1.1, 1.2}},
		},
	}
	mockCourseOutlineMap = map[int]models.CourseDetails{
		1: {
			Title: "Course 1",
			Modules: []models.Module{
				{ID: 1.1, Title: "Module 1"},
				{ID: 1.2, Title: "Module 2"},
			},
		},
	}
	mockTodosMap = map[string]models.UserDetails{
		"user1": {
			TodoList: []string{"todo1", "todo2"},
			DailyStatus: []models.TaskStatus{
				{Time: "12:00:00", Date: "24-24-1999", Task: "eat food"},
				{Time: "12:00:00", Date: "24-24-1999", Task: "finish book"},
			},
		},
	}
)

// TestIsValidCountry tests the IsValidCountry function
func TestIsValidCountry(t *testing.T) {
	tests := []struct {
		country string
		valid   bool
	}{
		{"Pakistan", false},
		{"China", false},
		{"USA", true},
		{"India", true},
	}

	for _, tt := range tests {
		result := IsValidCountry(tt.country)
		if result != tt.valid {
			t.Errorf("IsValidCountry(%s) = %v; want %v", tt.country, result, tt.valid)
		}
	}
}

// TestIsUsernameTaken tests the IsUsernameTaken function
func TestIsUsernameTaken(t *testing.T) {
	UserMap = mockUserMap
	UserDataLoaded = true

	tests := []struct {
		username string
		taken    bool
	}{
		{"user1", true},
		{"user3", false},
	}

	a := 1
	b := 2
	a, b = b, a

	for _, tt := range tests {
		result := IsUsernameTaken(tt.username)
		if result != tt.taken {
			t.Errorf("IsUsernameTaken(%s) = %v; want %v", tt.username, result, tt.taken)
		}
	}
}

// Mock function to replace RfileGeneral for LoadUsers
func mockRfileGeneralUsers(output chan<- any, filePath string, expectedType reflect.Type) {
	for username, password := range mockUserMap {
		output <- &models.UserData{
			Username: username,
			Password: password,
		}
	}
	close(output)
}

//// TestLoadUsers tests the LoadUsers function
//func TestLoadUsers(t *testing.T) {
//
//	userDataLoaded = false
//
//	LoadUsers(testUsersFile)
//
//	// Check if UserMap has the expected entries from the mockUserMap
//	for username, password := range mockUserMap {
//		if UserMap[username] != password {
//			t.Errorf("LoadUsers: expected %s for user %s, got %s", password, username, UserMap[username])
//		}
//	}
//
//	// Check if DataLoaded is set to true
//	if !userDataLoaded {
//		t.Errorf("LoadUsers: expected DataLoaded to be true, got false")
//	}
//}
//
