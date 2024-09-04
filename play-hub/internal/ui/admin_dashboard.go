package ui

import (
	"context"
	"fmt"
	"project2/internal/domain/entities"
	"strings"
)

func (ui *UI) ShowAdminDashboard() {
	for {
		fmt.Println("\033[1;35m") // Purple bold
		fmt.Println("╔═════════════════════════════════════╗")
		fmt.Println("║          🎮 Admin Dashboard 🎮      ║")
		fmt.Println("╚═════════════════════════════════════╝")
		fmt.Println("\033[0m") // Reset color

		fmt.Println("1. 🆕 Create a Game")
		fmt.Println("2. 🗑️ Delete a Game")
		fmt.Println("3. 📊 View User Stats")
		fmt.Println("4. 🚪 Logout")

		fmt.Print("\nEnter your choice: ")

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
			fmt.Println("\nLogging out... 👋")
			return
		default:
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Invalid choice. Please enter a number between 1 and 4.")
			fmt.Println("\033[0m") // Reset color
		}
	}
}

func (ui *UI) CreateGame() {
	var gameName string
	var maxPlayers int
	var minPlayers int
	var instances int

	fmt.Println("\033[1;34m") // Blue bold
	fmt.Println("\n🎮 Create a New Game")
	fmt.Println("\033[0m") // Reset color

	// Get game name
	for {
		fmt.Print("Enter the name of the game: ")
		gameName, _ = ui.reader.ReadString('\n')
		gameName = strings.TrimSpace(gameName)
		if gameName != "" {
			break
		} else {
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Game name cannot be empty. Please enter a valid name.")
			fmt.Println("\033[0m") // Reset color
		}
	}

	// Get maximum number of players
	for {
		fmt.Print("Enter the maximum number of players: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &maxPlayers)
		if err != nil || maxPlayers <= 0 {
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Invalid number of players. Please enter a positive integer.")
			fmt.Println("\033[0m") // Reset color
		} else {
			break
		}
	}

	// Get minimum number of players
	for {
		fmt.Print("Enter the minimum number of players: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &minPlayers)
		if err != nil || minPlayers <= 0 {
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Invalid number of players. Please enter a positive integer.")
			fmt.Println("\033[0m") // Reset color
		} else {
			break
		}
	}

	// Get instances of game available
	for {
		fmt.Print("Enter the number of instances: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &instances)
		if err != nil || instances <= 0 {
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Invalid number of players. Please enter a positive integer.")
			fmt.Println("\033[0m") // Reset color
		} else {
			break
		}
	}

	newGame := &entities.Game{
		GameName:   gameName,
		MaxPlayers: maxPlayers,
		MinPlayers: minPlayers,
		Instances:  instances,
	}
	_, err := ui.gameService.CreateGame(context.Background(), newGame)
	if err != nil {
		fmt.Printf("\033[1;31m❌ Error creating game: %v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32m") // Green bold
	fmt.Println("✅ Game created successfully!")
	fmt.Println("\033[0m") // Reset color
}

func (ui *UI) DeleteGame() {
	fmt.Println("\033[1;31m") // Red bold
	fmt.Println("\n🗑️ Deleting a Game...")
	fmt.Println("\033[0m") // Reset color

	// Fetch all games
	games, err := ui.gameService.GetAllGames(context.Background())
	if err != nil {
		fmt.Printf("\033[1;31m❌ Error retrieving games: %v\033[0m\n", err)
		return
	}

	// Check if there are any games
	if len(games) == 0 {
		fmt.Println("\033[1;33m⚠️ No games available to delete.\033[0m")
		return
	}

	// Display the list of games
	fmt.Println("\033[1;34mAvailable games:\033[0m")
	for i, game := range games {
		fmt.Printf("%d. %s (Max Players: %d)\n", i+1, game.GameName, game.MaxPlayers)
	}

	// Ask the user to choose a game to delete
	var choice int
	for {
		fmt.Print("Enter the number corresponding to the game you want to delete: ")
		input, _ := ui.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &choice)
		if err != nil || choice < 1 || choice > len(games) {
			fmt.Println("\033[1;31m❌ Invalid choice. Please enter a valid number.\033[0m")
		} else {
			break
		}
	}

	// Get the selected game ID and delete the game
	selectedGame := games[choice-1]
	err = ui.gameService.DeleteGame(context.Background(), selectedGame.GameID)
	if err != nil {
		fmt.Printf("\033[1;31m❌ Error deleting game: %v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32m") // Green bold
	fmt.Println("✅ Game deleted successfully!")
	fmt.Println("\033[0m") // Reset color
}

func (ui *UI) ViewUserStats() {
	fmt.Println("\033[1;34m") // Blue bold
	fmt.Println("\n📊 Viewing User Stats...")
	fmt.Println("\033[0m") // Reset color

	// Get the user's email ID
	var email string
	for {
		fmt.Print("Enter the email ID of the user: ")
		email, _ = ui.reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if email != "" {
			break
		} else {
			fmt.Println("\033[1;31m") // Red bold
			fmt.Println("❌ Email ID cannot be empty. Please enter a valid email ID.")
			fmt.Println("\033[0m") // Reset color
		}
	}

	// Retrieve user by email
	user, err := ui.userService.GetUserByEmail(context.Background(), email)
	if err != nil {
		fmt.Printf("\033[1;31m❌ Error retrieving user stats: %v\033[0m\n", err)
		return
	}

	// Display user stats
	fmt.Println("\033[1;32mUser Stats\033[0m")
	fmt.Printf("📧 Email: %s\n", user.Email)
	fmt.Printf("👤 Gender: %v\n", user.Gender)
	fmt.Printf("📞 Phone Number: %v\n", user.MobileNumber)
	fmt.Printf("🎖️ Role: %s\n", user.Role)
}
