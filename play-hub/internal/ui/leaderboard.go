package ui

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"project2/pkg/utils"
)

func (ui *UI) ViewLeaderboard() {
	fmt.Println("🏆 Leaderboard 🏆")
	users, err := ui.leaderboardService.GetOverallLeaderboard()
	if err != nil {
		fmt.Println("⚠️ Error fetching leaderboard:", err)
		return
	}

	// If there are no users on the leaderboard
	if len(users) == 0 {
		fmt.Println("😕 No users found on the leaderboard.")
		return
	}

	// Create a table for the leaderboard
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank 🥇", "Name 👤", "Score 💯"})

	// Iterate through users and add them to the table
	for i, user := range users {
		rank := i + 1
		name := utils.GetNameFromEmail(user.Email)
		score := user.OverallScore

		// Add the row to the table
		table.Append([]string{
			fmt.Sprintf("#%d", rank),
			name,
			fmt.Sprintf("%d", score),
		})
	}

	// Render the table to the console
	table.Render()

	fmt.Println("🏅 Keep playing to improve your rank!")
}
