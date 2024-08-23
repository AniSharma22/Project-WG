package course

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"project/internal/utils"
	"strconv"
	"strings"
)

// Function to get the number of decimal places in a float64
func countDecimalPlaces(f float64) int {
	// Convert the float to a string
	floatStr := strconv.FormatFloat(f, 'f', -1, 64)

	// Split the string on the decimal point
	parts := strings.Split(floatStr, ".")

	// If there is no decimal point, return 0
	if len(parts) < 2 {
		return 0
	}

	// Return the length of the part after the decimal point
	return len(parts[1])
}

// ManageCourses handles course-related actions
func ManageCourses(username string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("<---- Assigned Courses ---->")

		// Display assigned courses
		assignedCourses := utils.ProgressMap[username]
		for i, val := range assignedCourses {
			courseTitle := utils.CourseOutlineMap[val.CourseID].Title
			fmt.Printf("%d. %s (Course ID: %d)\n", i+1, courseTitle, val.CourseID)
		}

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

			if courseID <= 0 || courseID > len(utils.CourseOutlineMap) {
				fmt.Println("Invalid Course ID")
				break
			}

			// Check if the entered Course ID is among the assigned courses
			validCourseID := false
			for _, course := range assignedCourses {
				if course.CourseID == courseID {
					validCourseID = true
					break
				}
			}

			if !validCourseID {
				fmt.Println("The Course ID you entered is not assigned to you.")
				break
			}

			courseDetails, exists := utils.CourseOutlineMap[courseID]
			if !exists {
				fmt.Println("Course not found.")
				break
			}

			// Display modules and completed status
			counter := 0
			completedModules := make(map[float64]bool)
			for _, progress := range utils.ProgressMap[username] {
				if courseID == progress.CourseID {
					for _, section := range progress.CompletedSections {
						counter++
						completedModules[section] = true
					}
				}
			}
			fmt.Println("\033[1;32m") // Green bold
			fmt.Println("===================================================================")
			fmt.Println("\033[0m") // Reset color
			fmt.Println(counter, "/", len(utils.CourseOutlineMap[courseID].Modules), " Modules Completed", "        ", (float64(counter)/float64(len(utils.CourseOutlineMap[courseID].Modules)))*100, "% Completed")
			fmt.Printf("\nCourse Title: %s\n", courseDetails.Title)
			fmt.Println("Modules:")
			for i, module := range courseDetails.Modules {
				if completedModules[module.ID] {
					fmt.Printf("Module ID: %d.%d, Title: %s ✅\n", int(module.ID), i+1, module.Title)
				} else {
					fmt.Printf("Module ID: %d.%d, Title: %s \n", int(module.ID), i+1, module.Title)
				}
			}
			fmt.Println("\033[1;32m") // Green bold
			fmt.Println("===================================================================")
			fmt.Println("\033[0m") // Reset color
			// Mark module as completed
			fmt.Print("Enter the Module ID to mark as completed (or 0 to go back): ")
			moduleIDStr, _ := reader.ReadString('\n')
			moduleIDStr = strings.TrimSpace(moduleIDStr)

			if _, err := strconv.Atoi(moduleIDStr); err == nil {
				fmt.Println("Invalid Module ID. Please enter a valid number.")
				break
			}

			// Convert moduleIDStr to float64
			moduleID, err := strconv.ParseFloat(moduleIDStr, 64)
			if err != nil {
				fmt.Println("Invalid Module ID. Please enter a valid number.")
				break
			}
			fmt.Println(moduleID)
			numOfDecimalValues := countDecimalPlaces(moduleID)
			if numOfDecimalValues > 1 {
				fmt.Println("Invalid or Completed Module ID. Please enter a valid number.")
				break
			}

			if moduleID == 0 {
				break // Go back
			}
			bool1, bool2 := completedModules[moduleID]
			if bool1 && bool2 {
				fmt.Println("Module Already Marked as Completed!")
				break
			}
			var firstVal = int(moduleID)
			var _ = int((moduleID - math.Floor(moduleID)) * 10)

			if firstVal != courseID || moduleID < float64(courseID) || moduleID > float64(courseID+1) {
				fmt.Println("Invalid or Completed Module ID. Please enter a valid number.")
				break
			}

			if !completedModules[moduleID] {
				// Mark module as completed
				for i, progress := range utils.ProgressMap[username] {
					if progress.CourseID == courseID {
						utils.ProgressMap[username][i].CompletedSections = append(
							utils.ProgressMap[username][i].CompletedSections,
							moduleID,
						)
						for _, v := range utils.CourseOutlineMap[courseID].Modules {
							if v.ID == moduleID {
								_ = utils.AddToDailyStatus(username, v.Title)
							}
						}

						fmt.Println("Module marked as completed!")

						break
					}
				}
			} else {
				fmt.Println("Invalid Module ID")

			}

		case "2":
			return // Go back (exit the function)

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
