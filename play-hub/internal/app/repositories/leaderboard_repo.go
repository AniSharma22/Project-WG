package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
)

type leaderboardRepo struct {
	collection *mongo.Collection
}

func NewLeaderboardRepo(client *mongo.Client) interfaces.LeaderboardRepository {
	return &leaderboardRepo{
		collection: client.Database(config.DB.DBName).Collection(config.DB.LeaderboardsCollection),
	}
}

func (r *leaderboardRepo) GetGameLeaderboard(gameId primitive.ObjectID) ([]entities.Leaderboard, error) {
	filter := bson.M{"gameId": gameId}
	opts := options.Find().SetSort(bson.D{{"score", -1}}) // Sort by score in descending order

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var leaderboards []entities.Leaderboard
	if err := cursor.All(context.Background(), &leaderboards); err != nil {
		return nil, err
	}

	return leaderboards, nil
}

func (r *leaderboardRepo) GetOverallLeaderboard() ([]entities.Leaderboard, error) {
	opts := options.Find().SetSort(bson.D{{"score", -1}}) // Sort by score in descending order

	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var leaderboards []entities.Leaderboard
	if err := cursor.All(context.Background(), &leaderboards); err != nil {
		return nil, err
	}

	return leaderboards, nil
}

func (r *leaderboardRepo) AddOrUpdateLeaderboardEntry(entry *entities.Leaderboard) error {
	filter := bson.M{
		"gameId": entry.GameID,
		"userId": entry.UserID,
	}

	update := bson.M{
		"$set": bson.M{
			"score": entry.Score,
			"rank":  entry.Rank,
		},
	}

	upsertOpts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(context.Background(), filter, update, upsertOpts)
	return err
}
