package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type gameHistoryRepo struct {
	collection *mongo.Collection
}

func NewGameHistoryRepo() interfaces.GameHistoryRepository {
	return &gameHistoryRepo{
		collection: globals.Client.Database(config.DBName).Collection("GameHistory"),
	}
}

func (r *gameHistoryRepo) AddGameHistory(history *entities.GameHistory) error {
	_, err := r.collection.InsertOne(context.Background(), history)
	if err != nil {
		fmt.Println("Error inserting game history:", err)
		return err
	}
	return nil
}

func (r *gameHistoryRepo) RemoveGameHistory(historyID primitive.ObjectID) error {
	filter := bson.M{"_id": historyID}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Error deleting game history:", err)
		return err
	}
	return nil
}

func (r *gameHistoryRepo) FindGameHistoryByID(historyID primitive.ObjectID) (*entities.GameHistory, error) {
	filter := bson.M{"_id": historyID}
	var history entities.GameHistory
	err := r.collection.FindOne(context.Background(), filter).Decode(&history)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("game history not found")
		}
		fmt.Println("Error finding game history:", err)
		return nil, err
	}
	return &history, nil
}

func (r *gameHistoryRepo) GetAllGameHistories() ([]entities.GameHistory, error) {
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error finding game histories:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var histories []entities.GameHistory
	for cursor.Next(context.Background()) {
		var history entities.GameHistory
		if err := cursor.Decode(&history); err != nil {
			fmt.Println("Error decoding game history:", err)
			return nil, err
		}
		histories = append(histories, history)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return histories, nil
}

func (r *gameHistoryRepo) GetUserGameHistory(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	filter := bson.M{"userId": userId}
	var histories []entities.GameHistory

	// Execute the query to find matching game histories
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error finding game histories:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor to decode each document into a GameHistory object
	for cursor.Next(context.Background()) {
		var history entities.GameHistory
		if err := cursor.Decode(&history); err != nil {
			fmt.Println("Error decoding game history:", err)
			return nil, err
		}
		histories = append(histories, history)
	}

	// Check if there were any errors during the cursor iteration
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return histories, nil
}

func (r *gameHistoryRepo) GetResultsToUpdate(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	filter := bson.M{"userId": userId, "result": ""}
	var histories []entities.GameHistory

	// Execute the query to find matching game histories
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error finding game histories:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor to decode each document into a GameHistory object
	for cursor.Next(context.Background()) {
		var history entities.GameHistory
		if err := cursor.Decode(&history); err != nil {
			fmt.Println("Error decoding game history:", err)
			return nil, err
		}
		histories = append(histories, history)
	}

	// Check if there were any errors during the cursor iteration
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return histories, nil
}

func (r *gameHistoryRepo) GetCurrentDayHistory(userId primitive.ObjectID) ([]entities.GameHistory, error) {
	// Get the current date and the start and end of the day
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Construct the filter to match userId and createdAt within the current day
	filter := bson.M{
		"userId": userId,
		"createdAt": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	// Perform the query and get a cursor
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each document into the history slice
	var history []entities.GameHistory
	for cursor.Next(context.Background()) {
		var gameHistory entities.GameHistory
		if err := cursor.Decode(&gameHistory); err != nil {
			return nil, err
		}
		history = append(history, gameHistory)
	}

	// Check for any errors encountered during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return history, nil
}

func (r *gameHistoryRepo) UpdateResult(result string, slotId primitive.ObjectID, userID primitive.ObjectID) error {
	filter := bson.M{"userId": userID, "slotId": slotId}
	update := bson.M{
		"$set": bson.M{
			"result": result,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating game history:", err)
		return err
	}
	return nil
}
