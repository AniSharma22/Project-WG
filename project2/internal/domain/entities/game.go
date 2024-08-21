package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	MaxCapacity int                `bson:"maxCapacity"`
}
