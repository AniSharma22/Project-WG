package entities

import "time"

type Slot struct {
	SlotID       string    `json:"slot_id"`
	GameID       string    `json:"game_id"`
	Time         time.Time `json:"time"`
	BookedBy     User      `json:"booked_by"`
	InvitedUsers []User    `json:"invited_users"`
}
