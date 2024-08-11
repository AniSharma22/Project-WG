package dailyStatus

import (
	"fmt"
	"project/internal/utils"
)

// ViewDailyStatus shows the user's daily status
func ViewDailyStatus(username string) {

	fmt.Println("---------------------------------")
	fmt.Println("          DAILY STATUS ")
	fmt.Println("---------------------------------")
	// Retrieve the user's daily status
	userStatus, exists := utils.TodosMap[username]
	if !exists {
		fmt.Println("User not found.")
		return
	}

	// Check if the user has any daily status entries
	if len(userStatus.DailyStatus) == 0 {
		fmt.Println("No daily status entries found for the user.")
		return
	}

	// Display the daily status entries
	fmt.Printf("Date\t\t\tTime\t\tTask\n")
	for _, status := range userStatus.DailyStatus {
		fmt.Printf("%s\t\t%s\t\t%s\n", status.Date, status.Time, status.Task)

	}
	fmt.Println("---------------------------------")
	fmt.Println("---------------------------------")
}
