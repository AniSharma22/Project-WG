package repositories

import (
	"context"
	"errors"
	"fmt"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"project2/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *userRepo) AddToInvites(userId primitive.ObjectID, slotId primitive.ObjectID) error {
	// Create a filter to find the user by ID
	filter := bson.M{"_id": userId}

	// Create an update to add the invite to the InvitedSlots array
	update := bson.M{
		"$push": bson.M{
			"invitedSlots": slotId,
		},
	}

	// Perform the update operation
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating user with invited slot:", err)
		return err
	}

	return nil
}

func (r *userRepo) GetUserByEmail(email string) (*entities.User, error) {
	filter := bson.M{"email": email}

	var user entities.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found for email: %s", email)
		}
		return nil, fmt.Errorf("error querying database: %w", err)
	}

	if user.ID.IsZero() {
		return nil, fmt.Errorf("user found but ID is zero for email: %s", email)
	}

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
	// Retrieve the current user stats
	var user entities.User
	err := r.collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&user)
	if err != nil {
		fmt.Println("Error fetching user data:", err)
		return err
	}

	// Calculate new score
	newScore := utils.GetTotalScore(user.Wins+1, user.Losses)

	// Update the user's stats
	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "wins", Value: 1},
			{Key: "overallScore", Value: newScore - user.OverallScore},
		}},
	}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating user wins:", err)
		return err
	}
	return nil
}

func (r *userRepo) AddLoss(userId primitive.ObjectID) error {
	// Retrieve the current user stats
	var user entities.User
	err := r.collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&user)
	if err != nil {
		fmt.Println("Error fetching user data:", err)
		return err
	}

	// Calculate new score
	newScore := utils.GetTotalScore(user.Wins, user.Losses+1)

	// Update the user's stats
	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "losses", Value: 1},
			{Key: "overallScore", Value: newScore - user.OverallScore}, // Adjust overallScore based on new calculation
		}},
	}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
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

func (r *userRepo) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
	filter := bson.D{{"_id", userId}}
	var user entities.User
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found for id: %s", userId)

		}
	}
	return &user, nil
}

func (r *userRepo) GetPendingInvites(email string) ([]primitive.ObjectID, error) {
	filter := bson.M{"email": email}
	var user entities.User

	// Find the user document
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found for email: %s", email)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	// Return the invites
	return user.InvitedSlots, nil
}

func (r *userRepo) DeleteInvite(slotId primitive.ObjectID) error {
	email := globals.ActiveUser

	// Define the filter to find the user by email
	filter := bson.M{"email": email}

	// Define the update to pull the slotId from the invitedSlots array
	update := bson.M{
		"$pull": bson.M{
			"invitedSlots": slotId,
		},
	}

	// Perform the update operation
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetAllUsersByScore() ([]entities.User, error) {
	var users []entities.User

	// Define the sort filter to sort by OverallScore in descending order
	opts := options.Find().SetSort(bson.M{"overallScore": -1})

	// Perform the find operation with the sort options
	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each user
	for cursor.Next(context.Background()) {
		var user entities.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
