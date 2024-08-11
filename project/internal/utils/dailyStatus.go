package utils

import (
	"fmt"
	"project/internal/models"
	"time"
)

// AddToDailyStatus adds a new entry to the daily status of a user.
func AddToDailyStatus(username, doneTodo string) error {
	// Get the current time and date
	currentTime := time.Now().Format("15:04:05")   // Use format for time
	currentDate := time.Now().Format("2006-01-02") // Use format for date

	// Create a new TaskStatus
	taskStatus := models.TaskStatus{
		Time: currentTime,
		Date: currentDate,
		Task: doneTodo,
	}

	// Retrieve the user's current daily status
	userStatus, exists := TodosMap[username]
	if !exists {
		return fmt.Errorf("user not found")
	}

	// Append the new task status to the user's daily status
	userStatus.DailyStatus = append(userStatus.DailyStatus, taskStatus)

	// Update the user's daily status in TodosMap
	TodosMap[username] = userStatus

	return nil
}
