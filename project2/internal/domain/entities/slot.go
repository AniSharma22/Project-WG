package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Result struct {
	UserID primitive.ObjectID `bson:"userId"`
	Result string             `bson:"result"`
}

type Slot struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	GameID      primitive.ObjectID   `bson:"gameId"`
	Date        string               `bson:"date"`
	StartTime   string               `bson:"startTime"`
	EndTime     string               `bson:"endTime"`
	BookedUsers []primitive.ObjectID `bson:"bookedUsers"`
	Results     []Result             `bson:"results"`
}
