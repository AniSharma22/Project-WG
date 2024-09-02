package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func (ui *UI) ViewLeaderboard() {
	// Step 1: Fetch and display the list of all games
	games, err := ui.gameService.GetAllGames(context.Background())
	if err != nil {
		fmt.Println("⚠️ Error fetching games:", err)
		return
	}

	// If no games are available
	if len(games) == 0 {
		fmt.Println("😕 No games available.")
		return
	}

	fmt.Println("🎮 Available Games:")
	for i, game := range games {
		fmt.Printf("%d. %s\n", i+1, game.GameName)
	}

	// Step 2: Ask the user to select a game by its number
	fmt.Print("\nSelect a game by number(press 0 to go back): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		return
	}

	// Convert input to an integer
	gameIndex, err := strconv.Atoi(input)
	if err != nil || gameIndex < 1 || gameIndex > len(games) {
		fmt.Println("⚠️ Invalid selection.")
		return
	}

	// Get the selected game
	selectedGame := games[gameIndex-1]

	// Step 3: Fetch the leaderboard for the selected game
	fmt.Printf("\n🏆 Leaderboard for %s 🏆\n", selectedGame.GameName)
	users, err := ui.leaderboardService.GetGameLeaderboard(context.Background(), selectedGame.GameID)
	if err != nil {
		fmt.Println("⚠️ Error fetching leaderboard:", err)
		return
	}

	// If there are no users on the leaderboard
	if len(users) == 0 {
		fmt.Println("😕 No users found on the leaderboard.")
		return
	}

	// Step 4: Create a table for the leaderboard
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank 🥇", "Name 👤", "Score 💯"})

	// Iterate through users and add them to the table
	for i, user := range users {
		rank := i + 1
		name := user.UserName
		score := user.Score // Adjust this if your score field has a different name in the leaderboard for games

		// Add the row to the table
		table.Append([]string{
			fmt.Sprintf("#%d", rank),
			name,
			fmt.Sprintf("%.2f", score),
		})
	}

	// Render the table to the console
	table.Render()

	fmt.Println("🏅 Keep playing to improve your rank!")
}
