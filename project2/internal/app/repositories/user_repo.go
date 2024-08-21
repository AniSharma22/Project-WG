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
	"project2/pkg/utils"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo() interfaces.UserRepository {
	return &userRepo{
		collection: globals.Client.Database(config.DBName).Collection("Users"),
	}
}

func (r *userRepo) CreateUser(user *entities.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}
	return nil
}

func (r *userRepo) GetUserByEmail(email string) (*entities.User, error) {
	// Create a filter to find a document with the specified email
	filter := bson.M{"email": email}

	// Perform the query to find one document with the specified email
	var user entities.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// No document found, return a custom error indicating user not found
			return nil, errors.New("user not found")
		}
		// Error occurred while querying
		return nil, err
	}

	// Return the user and nil as the error
	return &user, nil
}

// EmailAlreadyExists checks if the email already exists in the database.
func (r *userRepo) EmailAlreadyExists(email string) error {
	// Create a filter to find a document with the specified email
	filter := bson.M{"email": email}

	// Perform the query to find one document with the specified email
	var result entities.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		// No document found, email does not exist
		return nil
	} else if err != nil {
		// Error occurred while querying
		return err
	}

	fmt.Println(result)

	// Email already exists
	return nil
}

func (r *userRepo) AddWin(userId primitive.ObjectID) error {
	filter := bson.D{{"_id", userId}}
	update := bson.D{
		{"$inc", bson.D{
			{"wins", 1},
			{"overallScore", utils.GetTotalScore(1, 0)}, // Pass wins and losses for total score calculation
		}},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating user wins:", err)
		return err
	}
	return nil
}

func (r *userRepo) AddLoss(userId primitive.ObjectID) error {
	filter := bson.D{{"_id", userId}}
	update := bson.D{
		{"$inc", bson.D{
			{"losses", 1},
			{"overallScore", utils.GetTotalScore(0, 1)}, // Pass wins and losses for total score calculation
		}},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating user losses:", err)
		return err
	}
	return nil
}

func (r *userRepo) GetAllUsers() ([]entities.User, error) {
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error finding users:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []entities.User
	for cursor.Next(context.Background()) {
		var user entities.User
		if err := cursor.Decode(&user); err != nil {
			fmt.Println("Error decoding user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return users, nil
}
