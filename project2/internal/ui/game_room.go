package ui

import (
	"fmt"
	"strconv"
	"strings"
)

func (ui *UI) ShowGameRoom() {
	games, err := ui.gameService.GetAllGames()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(games) == 0 {
		fmt.Println("No games available.")
		return
	}

	fmt.Println("Available Games:")
	for i, game := range games {
		fmt.Printf("%d. %s\n", i+1, game.Name) // Assuming Game object has a Name field
	}

	fmt.Print("Enter the number of the game you want to select: ")
	choice, err := ui.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	choice = strings.TrimSpace(choice)

	// Convert the user input into an index
	gameIndex, err := strconv.Atoi(choice)
	if err != nil || gameIndex < 1 || gameIndex > len(games) {
		fmt.Println("Invalid choice. Please enter a valid number.")
		return
	}

	selectedGame := games[gameIndex-1]
	ui.HandleGameSelection(selectedGame) // Call a function to handle the selected game
}

// HandleGameSelection processes the selected game
func (ui *UI) HandleGameSelection(game Game) { // Replace Game with the actual type of game object
	fmt.Printf("You selected: %s\n", game.Name)
	// Add your logic here to handle the selected game
}
