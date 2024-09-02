package entities

import (
	"github.com/google/uuid"
	"time"
)

type Slot struct {
	SlotID    uuid.UUID `json:"slot_id" db:"slot_id"`
	GameID    uuid.UUID `json:"game_id" db:"game_id"`
	Date      time.Time `json:"slot_date" db:"slot_date"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
	IsBooked  bool      `json:"is_booked" db:"is_booked"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
