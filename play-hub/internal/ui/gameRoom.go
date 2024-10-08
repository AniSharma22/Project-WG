package ui

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"project2/internal/domain/entities"
	"project2/pkg/globals"
	"strconv"
	"strings"
	"time"
)

func (ui *UI) ShowGameRoom() {
	// Retrieve all games from the game service
	games, err := ui.gameService.GetAllGames(context.Background())
	if err != nil {
		fmt.Println("❌ Error fetching games:", err)
		return
	}

	// Display the list of games to the user
	fmt.Println("🎮 Available Games:")
	for i, game := range games {
		fmt.Printf("%d. %s\n", i+1, game.GameName)
	}
	// Add an option to go back
	fmt.Printf("%d. 🔙 Go Back\n", len(games)+1)

	// Prompt the user to select a game or go back
	var choice int
	for {
		fmt.Print("👉 Select an option by entering the corresponding number: ")
		input, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			continue
		}

		// Trim the newline character and spaces
		input = strings.TrimSpace(input)

		// Convert the input to an integer
		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(games)+1 {
			fmt.Println("❗ Invalid input. Please enter a number corresponding to an option.")
		} else {
			break
		}
	}

	// Handle the user's choice
	if choice == len(games)+1 {
		return
	}

	// Get the selected game object
	selectedGame := games[choice-1]

	// Pass the selected game object to another function
	ui.HandleSelectedGame(&selectedGame)
}

// HandleSelectedGame displays all slots for the selected game and handles the user's selection.
func (ui *UI) HandleSelectedGame(game *entities.Game) {
	fmt.Printf("✔️ You selected: %s\n", game.GameName)

	// Retrieve all slots for the selected game
	slots, err := ui.slotService.GetCurrentDayGameSlots(context.Background(), game.GameID)
	if err != nil {
		fmt.Println("❌ Error fetching slots:", err)
		return
	}

	// Check if there are any available slots
	if len(slots) == 0 {
		fmt.Println("⚠️ No slots available for this game.")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"S.No", "Slot Timings"})

	// Display the list of available slots to the user
	fmt.Println("🕒 Available Slots:")
	for i, slot := range slots {
		if time.Now().After(slot.StartTime) {
			table.Append([]string{
				color.New(color.FgRed).Sprintf("#%d", i+1),
				color.New(color.FgRed).Sprintf("%s - %s", slot.StartTime.Format("03:04 PM"), slot.EndTime.Format("03:04 PM")),
			})
		} else {
			table.Append([]string{
				color.New(color.FgBlue).Sprintf("#%d", i+1),
				color.New(color.FgBlue).Sprintf("%s - %s", slot.StartTime.Format("03:04 PM"), slot.EndTime.Format("03:04 PM")),
			})
		}
	}

	table.Render()

	// Prompt the user to select a slot
	var choice int
	for {
		fmt.Print("👉 Select a slot by entering the corresponding number: ")
		input, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			continue
		}

		// Trim the newline character and spaces
		input = strings.TrimSpace(input)

		// Convert the input to an integer
		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(slots) {
			fmt.Println("❗ Invalid input. Please enter a number corresponding to a slot.")
		} else {
			break
		}
	}

	// Get the selected slot object
	selectedSlot := slots[choice-1]

	// Pass the selected game and slot objects to another function
	ui.HandleSelectedSlot(game, &selectedSlot)
}

// HandleSelectedSlot processes the selected game and slot entities.
func (ui *UI) HandleSelectedSlot(game *entities.Game, slot *entities.Slot) {
	bookedUsers, _ := ui.bookingService.GetSlotBookedUsers(context.Background(), slot.SlotID)
	// Display the selected slot's time and game name
	fmt.Printf("\n📅 Slot Details:\n")
	fmt.Printf("🎮 Game: %s\n", game.GameName)
	fmt.Printf("⏰ Slot Time: %s to %s\n", slot.StartTime.Format("03:04 PM"), slot.EndTime.Format("03:04 PM"))

	// Display booked users
	if len(bookedUsers) == 0 {
		fmt.Println("🚫 Booked Users: None")
	} else {
		fmt.Println("👥 Booked Users:")
		for _, userName := range bookedUsers {
			fmt.Printf("- %s\n", userName)
		}
	}

	// Display results (winners and losers)
	//if len(slot.Results) == 0 {
	//	fmt.Println("⚖️ Results: No results recorded for this slot yet.")
	//} else {
	//	fmt.Println("🏅 Results:")
	//	for _, result := range slot.Results {
	//		user, _ := ui.userService.GetUserById(result.UserID)
	//		if result.Result == "win" {
	//			fmt.Printf("- 🏆 %s (User ID: %s) winner\n", utils.GetNameFromEmail(user.Email), result.UserID)
	//		} else {
	//			fmt.Printf("-   %s (User ID: %s) loser\n", utils.GetNameFromEmail(user.Email), result.UserID)
	//		}
	//	}
	//}

	// Show options to the user
	fmt.Println("\n🔧 Options:")
	fmt.Println("1. ✅ Book in this slot")
	fmt.Println("2. ✉️ Invite to this slot")
	fmt.Println("3. 🔙 Go back")

	// Handle user input
	var choice int
	for {
		fmt.Print("👉 Select an option by entering the corresponding number: ")
		input, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			continue
		}

		// Trim the newline character and spaces
		input = strings.TrimSpace(input)

		// Convert the input to an integer
		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("❗ Invalid input. Please enter a number between 1 and 3.")
		} else {
			break
		}
	}

	// Process user choice
	switch choice {
	case 1:
		err := ui.bookingService.MakeBooking(context.Background(), globals.ActiveUser, slot.SlotID)
		if err != nil {
			fmt.Println("❌", err)
			return
		}
		fmt.Println("🎉 Slot booked successfully!")
	case 2:
		fmt.Print("✉️ Enter the email of the user you want to invite to the slot: ")
		email, err := ui.reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading email:", err)
			return
		}

		// Trim the newline character from the email
		email = strings.TrimSpace(email)

		// Assuming you have a method to find the user by email
		user, err := ui.userService.GetUserByEmail(context.Background(), email)
		if err != nil {
			fmt.Println("❌ User not found or error retrieving user:", err)
			return
		}

		// Now pass the game, slot, and user ID to the InviteToSlot method
		_, err = ui.invitationService.MakeInvitation(context.Background(), globals.ActiveUser, user.UserID, slot.SlotID)
		if err != nil {
			fmt.Println("❌ Error inviting user to slot:", err)
			return
		}

		fmt.Println("✉️ User invited to slot successfully!")

	case 3:
		ui.ShowGameRoom()
	}
}
