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

type gameRepo struct {
	collection *mongo.Collection
}

func NewGameRepo() interfaces.GameRepository {
	return &gameRepo{
		collection: globals.Client.Database(config.DBName).Collection("Games"),
	}
}

func (g *gameRepo) GetGameByID(gameId primitive.ObjectID) (*entities.Game, error) {
	var game entities.Game
	filter := bson.D{{"_id", gameId}}
	err := g.collection.FindOne(context.Background(), filter).Decode(&game)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("game not found")
		}
		fmt.Println("Error finding game:", err)
		return nil, err
	}
	return &game, nil
}

func (g *gameRepo) GetAllGames() ([]entities.Game, error) {
	var games []entities.Game
	cursor, err := g.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error finding games:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var game entities.Game
		if err := cursor.Decode(&game); err != nil {
			fmt.Println("Error decoding game:", err)
			return nil, err
		}
		games = append(games, game)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return games, nil
}

func (g *gameRepo) CreateGame(game *entities.Game) error {
	_, err := g.collection.InsertOne(context.Background(), game)
	if err != nil {
		fmt.Println("Error inserting game:", err)
		return err
	}
	return nil
}

func (g *gameRepo) DeleteGame(gameId primitive.ObjectID) error {
	_, err := g.collection.DeleteOne(context.Background(), bson.D{{"_id", gameId}})
	if err != nil {
		fmt.Println("Error deleting game:", err)
		return err
	}
	return nil
}
