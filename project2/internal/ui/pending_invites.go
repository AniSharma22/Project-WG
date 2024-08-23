package ui

import (
	"fmt"
	"project2/pkg/globals"
	"strconv"
	"strings"
)

func (ui *UI) ViewPendingInvites() {
	fmt.Println("Pending Invites")

	// Retrieve pending invites for the active user
	invites, err := ui.userService.GetPendingInvites(globals.ActiveUser)
	if err != nil {
		fmt.Println("Error retrieving pending invites:", err)
		return
	}

	// If there are no pending invites, inform the user and return
	if len(invites) == 0 {
		fmt.Println("No pending invites")
		return
	}

	// Display the list of pending invites with corresponding numbers
	for i, invite := range invites {
		fmt.Printf("%d. Slot ID: %s, Game ID: %s, Date: %s, Start Time: %s, End Time: %s\n",
			i+1, invite.SlotID.Hex(), invite.GameID.Hex(), invite.Date, invite.StartTime, invite.EndTime)
	}

	// Ask the user to choose an invite by number
	fmt.Print("Enter the number of the invite you want to respond to: ")
	choiceStr, _ := ui.reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(invites) {
		fmt.Println("Invalid choice. Please enter a valid number.")
		return
	}

	// Get the selected invite based on the user's input
	selectedInvite := invites[choice-1]

	// Ask the user to accept or reject the selected invite
	fmt.Println("Do you want to accept or reject this invite? (a/r):")
	input, _ := ui.reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "a":
		err := ui.userService.AcceptInvite(selectedInvite)
		if err != nil {
			fmt.Printf("Error accepting invite #%d: %v\n", choice, err)
		} else {
			fmt.Printf("Invite #%d accepted\n", choice)
		}
	case "r":
		err := ui.userService.RejectInvite(selectedInvite)
		if err != nil {
			fmt.Printf("Error rejecting invite #%d: %v\n", choice, err)
		} else {
			fmt.Printf("Invite #%d rejected\n", choice)
		}
	default:
		fmt.Println("Invalid option, please choose 'a' to accept or 'r' to reject.")
	}
}
