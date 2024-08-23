package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SlotID    primitive.ObjectID `bson:"slotId"`
	UserID    primitive.ObjectID `bson:"userId"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"createdAt"`
}
