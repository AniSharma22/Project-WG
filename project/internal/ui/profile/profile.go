package profile

import (
	"fmt"
	"project/internal/utils"
)

// ViewProfile displays the user's profile
func ViewProfile(username string) {
	// Print the username
	fmt.Println("---------------------------------")
	fmt.Println("           PROFILE")
	fmt.Println("---------------------------------")
	fmt.Printf("Username: %s\n", username)

	// Retrieve and print the assigned courses
	assignedCourses, exists := utils.ProgressMap[username]
	if !exists || len(assignedCourses) == 0 {
		fmt.Println("No assigned courses found for this user.")
		return
	}

	fmt.Println("Assigned Courses:")
	for _, progress := range assignedCourses {
		courseTitle := utils.CourseOutlineMap[progress.CourseID].Title
		fmt.Printf("- %s (Course ID: %d)\n", courseTitle, progress.CourseID)

	}
	fmt.Println("---------------------------------")
	fmt.Println("---------------------------------")
}
