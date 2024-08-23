package ui

import (
	"fmt"
	"project2/pkg/globals"
)

func (ui *UI) ViewPendingInvites() {
	fmt.Println("Pending Invites")

	invites, err := ui.userService.GetPendingInvites(globals.ActiveUser)
	if err != nil {
		fmt.Println("Error retrieving pending invites:", err)
		return
	}
	if len(invites) == 0 {
		fmt.Println("No pending invites")
		return
	}

	for i, invite := range invites {
		fmt.Printf("Invite #%d:\n", i+1)
		fmt.Printf("Slot ID: %s\n", invite.SlotID.Hex())
		fmt.Printf("Game ID: %s\n", invite.GameID.Hex())
		fmt.Printf("Date: %s\n", invite.Date)
		fmt.Printf("Start Time: %s\n", invite.StartTime)
		fmt.Printf("End Time: %s\n", invite.EndTime)
		fmt.Println("Do you want to accept or reject this invite? (a/r):")

		input, _ := ui.reader.ReadString('\n')
		input = input[:len(input)-1] // Remove newline character

		switch input {
		case "a":
			err := ui.userService.AcceptInvite(invite.SlotID)
			if err != nil {
				fmt.Printf("Error accepting invite #%d: %v\n", i+1, err)
			} else {
				fmt.Printf("Invite #%d accepted\n", i+1)
			}
		case "r":
			err := ui.userService.RejectInvite(invite.SlotID)
			if err != nil {
				fmt.Printf("Error rejecting invite #%d: %v\n", i+1, err)
			} else {
				fmt.Printf("Invite #%d rejected\n", i+1)
			}
		default:
			fmt.Println("Invalid option, please choose 'a' to accept or 'r' to reject.")
		}
	}
}
