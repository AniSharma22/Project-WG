package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
)

type slotRepo struct {
	collection *mongo.Collection
}

func NewSlotRepo() interfaces.SlotRepository {
	return &slotRepo{
		collection: globals.Client.Database(config.DBName).Collection("Slots"),
	}
}

// GetSlotsByDate retrieves all slots for a given date and game.
func (r *slotRepo) GetSlotsByDate(date string, gameId primitive.ObjectID) ([]entities.Slot, error) {
	filter := bson.M{"date": date, "gameId": gameId}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve slots: %w", err)
	}
	defer cursor.Close(context.Background())

	var slots []entities.Slot
	if err := cursor.All(context.Background(), &slots); err != nil {
		return nil, fmt.Errorf("failed to decode slots: %w", err)
	}

	return slots, nil
}

// GetSlotByDateAndTime retrieves a specific slot by date and time.
func (r *slotRepo) GetSlotByDateAndTime(date string, gameId primitive.ObjectID, startTime string) (*entities.Slot, error) {
	filter := bson.M{"date": date, "gameId": gameId, "startTime": startTime}
	var slot entities.Slot
	err := r.collection.FindOne(context.Background(), filter).Decode(&slot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("this slot data is not available")
		}
		return nil, err
	}
	return &slot, nil
}

// BookSlot books a slot for a user.
func (r *slotRepo) BookSlot(userId primitive.ObjectID, date string, gameId primitive.ObjectID, startTime string) error {

	filter := bson.M{"date": date, "gameId": gameId, "startTime": startTime}
	update := bson.M{"$addToSet": bson.M{"bookedUsers": userId}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *slotRepo) InsertSlot(slot entities.Slot) (*mongo.InsertOneResult, error) {

	return r.collection.InsertOne(context.Background(), slot)
}
func (r *slotRepo) GetSlotById(slotId primitive.ObjectID) (*entities.Slot, error) {
	filter := bson.M{"_id": slotId}
	var slot entities.Slot
	err := r.collection.FindOne(context.Background(), filter).Decode(&slot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("this slot data is not available")

		}
		return nil, err
	}
	return &slot, nil
}
