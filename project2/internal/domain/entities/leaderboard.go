package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Leaderboard struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	GameID primitive.ObjectID `bson:"gameId"`
	UserID primitive.ObjectID `bson:"userId"`
	Score  float64            `bson:"score"`
	Rank   int                `bson:"rank"`
}
