package models

import (
	"github.com/google/uuid"
	"time"
)

type Invitations struct {
	InvitationId uuid.UUID
	SlotId       uuid.UUID
	GameName     string
	Date         time.Time
	StartTime    time.Time
	EndTime      time.Time
	BookedUsers  []string
	InvitedBy    string
}

type Bookings struct {
	BookingId   uuid.UUID
	GameName    string
	Date        time.Time
	StartTime   time.Time
	EndTime     time.Time
	BookedUsers []string
}

type Leaderboard struct {
	UserName string
	Score    float64
}
