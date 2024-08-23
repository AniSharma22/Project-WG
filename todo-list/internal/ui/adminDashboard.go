package ui

import (
	"bufio"
	"fmt"
	"os"
	"project/internal/models"
	"project/internal/utils"
	"strconv"
	"strings"
)

// HandleAdminDashboard provides options for the admin to manage users and courses
func HandleAdminDashboard(username string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n<- Admin Dashboard ->")
		fmt.Println("1. View a particular user")
		fmt.Println("2. View all courses")
		fmt.Println("3. Logout")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleViewUser(reader)
		case "2":
			handleViewAllCourses(reader)
		case "3":
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// handleViewUser handles the logic to view a particular user's details
func handleViewUser(reader *bufio.Reader) {
	fmt.Print("Enter the username of the user: ")
	userToView, _ := reader.ReadString('\n')
	userToView = strings.TrimSpace(userToView)

	userDetails, exists := utils.ProgressMap[userToView]
	if !exists {
		fmt.Println("User not found.")
		return
	}
	fmt.Printf("\n\n")
	assignedCoursesSlice := make([]bool, len(utils.CourseOutlineMap)+1)
	fmt.Printf("User: %s\n", userToView)
	fmt.Println("------------------------------")
	fmt.Println("assigned courses:")
	for _, progress := range userDetails {
		assignedCoursesSlice[progress.CourseID] = true
		courseTitle := utils.CourseOutlineMap[progress.CourseID].Title
		fmt.Printf("Course: %s (Course ID: %d)\n", courseTitle, progress.CourseID)
		fmt.Println("Completed Sections:", progress.CompletedSections)
	}
	fmt.Println("------------------------------")
	fmt.Println("Unassigned courses")
	for i, v := range assignedCoursesSlice {
		if i != 0 && !v == true {
			fmt.Printf("Course: %s (Course ID: %d)\n", utils.CourseOutlineMap[i].Title, i)
		}
	}
	fmt.Println("------------------------------")

	fmt.Println("\nOptions:")
	fmt.Println("1. Assign a new course")
	fmt.Println("2. Deassign a course")
	fmt.Println("3. Go back")

	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		assignCourse(reader, userToView)
	case "2":
		deassignCourse(reader, userToView)
	case "3":
		return // Go back
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}

}

// handleViewAllCourses handles the logic to view all courses
func handleViewAllCourses(reader *bufio.Reader) {
	//fmt.Println(utils.CourseOutlineMap)
	fmt.Printf("\n")
	fmt.Println("------------------------------")
	fmt.Println("All Courses:")
	for courseID, course := range utils.CourseOutlineMap {
		fmt.Printf("Course ID: %d, Title: %s\n", courseID, course.Title)
	}
	fmt.Println("------------------------------")
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. View details of a particular course")
		fmt.Println("2. Go back")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter the Course ID to view details: ")
			courseIDStr, _ := reader.ReadString('\n')
			courseIDStr = strings.TrimSpace(courseIDStr)

			// Validate and convert Course ID
			courseID, err := strconv.Atoi(courseIDStr)
			if err != nil {
				fmt.Println("Invalid Course ID. Please enter a valid number.")
				break
			}
			if courseID < 1 || courseID > len(utils.CourseOutlineMap) {
				fmt.Println("Invalid Course ID. Please enter a valid number.")
				break
			}
			// Display course details
			courseDetails, exists := utils.CourseOutlineMap[courseID]
			if !exists {
				fmt.Println("Course not found.")
				break
			}
			fmt.Println("------------------------------")
			fmt.Printf("Course Title: %s\n", courseDetails.Title)
			fmt.Println("Modules:")
			for _, module := range courseDetails.Modules {
				fmt.Printf("  Module ID: %.1f, Title: %s\n", module.ID, module.Title)
			}
			fmt.Println("------------------------------")

		case "2":
			return // Go back

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// assignCourse handles the logic to assign a new course to a user
func assignCourse(reader *bufio.Reader, viewedUser string) {
	fmt.Print("Enter the Course ID to assign: ")
	courseIDStr, _ := reader.ReadString('\n')
	courseIDStr = strings.TrimSpace(courseIDStr)

	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil || courseID <= 0 || courseID > len(utils.CourseOutlineMap) {
		fmt.Println("Invalid Course ID. Please enter a valid number.")
		return
	}

	for _, progress := range utils.ProgressMap[viewedUser] {
		if progress.CourseID == courseID {
			fmt.Println("Course is already assigned to the user.")
			return
		}
	}

	newProgress := utils.ProgressMap[viewedUser]
	newCourse := models.CompletedSection{
		CourseID:          courseID,
		CompletedSections: []float64{},
	}
	newProgress = append(newProgress, newCourse)
	utils.ProgressMap[viewedUser] = newProgress
	fmt.Println("Course assigned successfully!")
}

// deassignCourse handles the logic to deassign a course from a user
func deassignCourse(reader *bufio.Reader, viewedUser string) {
	fmt.Print("Enter the Course ID to deassign: ")
	courseIDStr, _ := reader.ReadString('\n')
	courseIDStr = strings.TrimSpace(courseIDStr)

	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil || courseID <= 0 || courseID > len(utils.CourseOutlineMap) {
		fmt.Println("Invalid Course ID. Please enter a valid number.")
		return
	}

	progressIndex := -1
	for i, v := range utils.ProgressMap[viewedUser] {
		if v.CourseID == courseID {
			progressIndex = i
		}
	}

	if progressIndex == -1 {
		fmt.Println("Course is not assigned to the user.")
		return
	}
	newProgress := utils.ProgressMap[viewedUser]
	newProgress = append(newProgress[:progressIndex], newProgress[progressIndex+1:]...)
	utils.ProgressMap[viewedUser] = newProgress
	fmt.Println("Course deassigned successfully!")
}
