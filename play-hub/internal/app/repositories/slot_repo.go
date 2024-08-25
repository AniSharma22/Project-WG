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
	"time"
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
func (r *slotRepo) GetSlotsByDate(date time.Time, gameId primitive.ObjectID) ([]entities.Slot, error) {
	date = date.UTC()
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
	// changing UTC TO IST before returning
	for i := range slots {
		slots[i].StartTime = slots[i].StartTime.In(globals.IstLocation)
		slots[i].EndTime = slots[i].EndTime.In(globals.IstLocation)
		slots[i].Date = slots[i].Date.In(globals.IstLocation)
	}
	return slots, nil
}

// GetSlotByDateAndTime retrieves a specific slot by date and time.
func (r *slotRepo) GetSlotByDateAndTime(date time.Time, gameId primitive.ObjectID, startTime time.Time) (*entities.Slot, error) {
	date = date.UTC()
	startTime = startTime.UTC()
	filter := bson.M{"date": date, "gameId": gameId, "startTime": startTime}
	var slot entities.Slot
	err := r.collection.FindOne(context.Background(), filter).Decode(&slot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("this slot data is not available")
		}
		return nil, err
	}
	slot.StartTime = slot.StartTime.In(globals.IstLocation)
	slot.EndTime = slot.EndTime.In(globals.IstLocation)
	slot.Date = slot.Date.In(globals.IstLocation)
	return &slot, nil
}

// BookSlot books a slot for a user.
func (r *slotRepo) BookSlot(userId primitive.ObjectID, slotId primitive.ObjectID) error {
	// Define the filter to find the slot by its ID
	filter := bson.M{"_id": slotId}

	// Define the update to add the userId to the BookedUsers slice
	update := bson.M{
		"$addToSet": bson.M{
			"bookedUsers": userId,
		},
	}

	// Perform the update operation
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error booking slot: %w", err)
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
	slot.StartTime = slot.StartTime.In(globals.IstLocation)
	slot.EndTime = slot.EndTime.In(globals.IstLocation)
	slot.Date = slot.Date.In(globals.IstLocation)
	return &slot, nil
}

func (r *slotRepo) GetUpcomingBookedSlots(userId primitive.ObjectID) ([]entities.Slot, error) {
	today := time.Now().Truncate(24 * time.Hour).UTC() // Truncate to get the date without the time part
	currentTime := time.Now().UTC()

	filter := bson.M{
		"date": today,
		"startTime": bson.M{
			"$gte": currentTime,
		},
		"bookedUsers": userId,
	}

	// Find all matching slots
	var slots []entities.Slot
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	// Decode the results into the slots slice
	if err = cursor.All(context.Background(), &slots); err != nil {
		return nil, err
	}

	for _, slot := range slots {
		slot.StartTime = slot.StartTime.In(globals.IstLocation)
		slot.EndTime = slot.EndTime.In(globals.IstLocation)
		slot.Date = slot.Date.In(globals.IstLocation)
	}

	return slots, nil
}

func (r *slotRepo) AddResultToSlot(userId primitive.ObjectID, slotId primitive.ObjectID, result string) error {
	resultToAdd := entities.Result{
		UserID: userId,
		Result: result,
	}

	filter := bson.M{"_id": slotId}
	update := bson.M{
		"$addToSet": bson.M{
			"results": resultToAdd,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
