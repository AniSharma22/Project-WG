package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invite struct {
	SlotId      primitive.ObjectID
	GameName    string
	Date        time.Time
	StartTime   time.Time
	EndTime     time.Time
	BookedUsers []string
}
