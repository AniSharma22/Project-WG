package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Email        string               `bson:"email"`
	Password     string               `bson:"password"`
	PhoneNo      string               `bson:"phoneNumber"`
	Gender       string               `bson:"gender"`
	Wins         int                  `bson:"wins"`
	Losses       int                  `bson:"losses"`
	OverallScore float32              `bson:"overallScore"`
	InvitedSlots []primitive.ObjectID `bson:"invitedSlots"`
	Role         string               `bson:"role"`
}
