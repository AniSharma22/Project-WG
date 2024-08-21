package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GameHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId"`
	GameID    primitive.ObjectID `bson:"gameId"`
	SlotID    primitive.ObjectID `bson:"slotId"`
	Result    string             `bson:"result"`
	CreatedAt time.Time          `bson:"createdAt"`
}
