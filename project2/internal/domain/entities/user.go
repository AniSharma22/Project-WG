package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	PhoneNo      string             `bson:"phoneNumber"`
	Gender       string             `bson:"gender"`
	Wins         int                `bson:"wins"`
	Losses       int                `bson:"losses"`
	OverallScore int                `bson:"overallScore"`
	InvitedSlots []InvitedSlot      `bson:"invitedSlots"`
	Role         string             `bson:"role"`
}

type InvitedSlot struct {
	SlotID    primitive.ObjectID `bson:"slotId"`
	GameID    primitive.ObjectID `bson:"gameId"`
	Date      time.Time          `bson:"date"`
	StartTime time.Time          `bson:"startTime"`
	EndTime   time.Time          `bson:"endTime"`
}
