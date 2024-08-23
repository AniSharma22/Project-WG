package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
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
	Date      string             `bson:"date"`
	StartTime string             `bson:"startTime"`
	EndTime   string             `bson:"endTime"`
}
