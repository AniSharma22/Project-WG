package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Result struct {
	UserID primitive.ObjectID `bson:"userId"`
	Result string             `bson:"result"`
}

type Slot struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	GameID      primitive.ObjectID   `bson:"gameId"`
	Date        time.Time            `bson:"date"`
	StartTime   time.Time            `bson:"startTime"`
	EndTime     time.Time            `bson:"endTime"`
	BookedUsers []primitive.ObjectID `bson:"bookedUsers"`
	Results     []Result             `bson:"results"`
}
