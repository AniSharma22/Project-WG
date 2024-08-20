package entities

import "time"

type Slot struct {
	Date  string      `json:"date"`  // Date of the slots
	Games []GameSlots `json:"games"` // Slots grouped by games
}

type GameSlots struct {
	GameID string      `json:"game_id"` // ID of the game
	Slots  []SlotStats `json:"slots"`   // List of slots for the game
}

type SlotStats struct {
	SlotID       string        `json:"slot_id"`
	Time         time.Time     `json:"time"`          // Slot start time
	BookedBy     []User        `json:"booked_by"`     // Users who booked the slot
	InvitedUsers []User        `json:"invited_users"` // Users invited to the slot
	Duration     time.Duration `json:"duration"`      // Duration (should be 20 mins by default)
	IsBooked     bool          `json:"is_booked"`     // To indicate if the slot is booked
}
