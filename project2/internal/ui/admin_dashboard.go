package ui

import (
	"bufio"
	"fmt"
	"strings"
)

func (ui *UI) ShowAdminDashboard() {
	for {
		fmt.Println("\nAdmin Dashboard")
		fmt.Println("1. Create a Game")
		fmt.Println("2. Delete a Game")
		fmt.Println("3. View User Stats")
		fmt.Println("4. Logout")

		fmt.Print("Enter your choice: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			ui.CreateGame()
		case "2":
			ui.DeleteGame()
		case "3":
			ui.ViewUserStats()
		case "4":
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func (ui *UI) CreateGame() {
	var gameName string
	var maxPlayers int

	// Get game name
	for {
		fmt.Print("Enter the name of the game: ")
		gameName, _ = ui.reader.ReadString('\n')
		gameName = strings.TrimSpace(gameName)
		if gameName != "" {
			break
		} else {
			fmt.Println("Game name cannot be empty. Please enter a valid name.")
		}
	}

	// Get maximum number of players
	for {
		fmt.Print("Enter the maximum number of players: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &maxPlayers)
		if err != nil || maxPlayers <= 0 {
			fmt.Println("Invalid number of players. Please enter a positive integer.")
		} else {
			break
		}
	}

	err := ui.gameService.CreateGame(gameName, maxPlayers)
	if err != nil {
		fmt.Printf("Error creating game: %v\n", err)
		return
	}

	fmt.Println("Game created successfully!")
}

func (ui *UI) DeleteGame() {
	fmt.Println("Deleting a game...")

	// Fetch all games
	games, err := ui.gameService.GetAllGames()
	if err != nil {
		fmt.Printf("Error retrieving games: %v\n", err)
		return
	}

	// Check if there are any games
	if len(games) == 0 {
		fmt.Println("No games available to delete.")
		return
	}

	// Display the list of games
	fmt.Println("Available games:")
	for i, game := range games {
		fmt.Printf("%d. %s (Max Players: %d)\n", i+1, game.Name, game.MaxCapacity)
	}

	// Ask the user to choose a game to delete
	reader := bufio.NewReader(ui.reader)
	var choice int
	for {
		fmt.Print("Enter the number corresponding to the game you want to delete: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &choice)
		if err != nil || choice < 1 || choice > len(games) {
			fmt.Println("Invalid choice. Please enter a valid number.")
		} else {
			break
		}
	}

	// Get the selected game ID and delete the game
	selectedGame := games[choice-1]
	err = ui.gameService.DeleteGame(selectedGame.ID)
	if err != nil {
		fmt.Printf("Error deleting game: %v\n", err)
		return
	}

	fmt.Println("Game deleted successfully!")
}

func (ui *UI) ViewUserStats() {
	fmt.Println("Viewing user stats...")

	// Get the user's email ID
	var email string
	for {
		fmt.Print("Enter the email ID of the user: ")
		email, _ = ui.reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if email != "" {
			break
		} else {
			fmt.Println("Email ID cannot be empty. Please enter a valid email ID.")
		}
	}

	// Retrieve user by email
	user, err := ui.userService.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("Error retrieving user stats: %v\n", err)
		return
	}

	// Display user stats (assuming user has some stats fields to display)
	fmt.Printf("User Stats")
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Gender: %v\n", user.Gender)
	fmt.Printf("Phone Number: %v\n", user.PhoneNo)
	fmt.Printf("Role: %s\n", user.Role)
	fmt.Printf("Games Played: %d\n", user.Wins+user.Losses)
	fmt.Printf("Score: %d\n", user.OverallScore)

}
