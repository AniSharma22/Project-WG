package ui

import (
	"fmt"
	"strconv"
	"strings"
)

func (ui *UI) ViewPendingInvites() {
	fmt.Println("📬  Pending Invites  📬")

	// Retrieve pending invites for the active user
	invites, err := ui.userService.GetPendingInvites()
	if err != nil {
		fmt.Println("⚠️ Error retrieving pending invites:", err)
		return
	}

	// If there are no pending invites, inform the user and return
	if len(invites) == 0 {
		fmt.Println("🎉 No pending invites")
		return
	}

	// Display the list of pending invites with corresponding numbers
	fmt.Println("You have the following pending invites:")
	for i, invite := range invites {
		fmt.Printf(" %d️⃣  %s\n", i+1, invite.GameName)
		fmt.Printf("   🗓️  Date: %s\n", invite.Date)
		fmt.Printf("   🕒  Time: %s - %s\n", invite.StartTime, invite.EndTime)
		fmt.Println("   👥  Participants:")
		for _, user := range invite.BookedUsers {
			fmt.Printf("    - %s\n", user)
		}
		fmt.Println()
	}

	// Ask the user to choose an invite by number
	fmt.Print("Enter the number of the invite you want to respond to: ")
	choiceStr, _ := ui.reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(invites) {
		fmt.Println("❌ Invalid choice. Please enter a valid number.")
		return
	}

	// Get the selected invite based on the user's input
	selectedInvite := invites[choice-1]

	// Ask the user to accept or reject the selected invite
	fmt.Println("Would you like to accept or reject this invite? (a/r):")
	input, _ := ui.reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "a":
		err := ui.userService.AcceptInvite(selectedInvite.SlotId)
		if err != nil {
			fmt.Printf("❌ Error accepting invite #%d: %v\n", choice, err)
		} else {
			fmt.Printf("✅ Invite #%d accepted\n", choice)
		}
	case "r":
		err := ui.userService.RejectInvite(selectedInvite.SlotId)
		if err != nil {
			fmt.Printf("❌ Error rejecting invite #%d: %v\n", choice, err)
		} else {
			fmt.Printf("✅ Invite #%d rejected\n", choice)
		}
	default:
		fmt.Println("❌ Invalid option. Please choose 'a' to accept or 'r' to reject.")
	}
}
