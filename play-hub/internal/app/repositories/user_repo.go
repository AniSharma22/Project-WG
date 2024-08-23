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

func (r *userRepo) AddToInvites(userId primitive.ObjectID, invite entities.InvitedSlot) error {
	// Create a filter to find the user by ID
	fmt.Println("idhar aagya")
	fmt.Println(userId)
	filter := bson.M{"_id": userId}

	// Create an update to add the invite to the InvitedSlots array
	update := bson.M{
		"$push": bson.M{
			"invitedSlots": invite,
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

func (r *userRepo) GetPendingInvites(email string) ([]entities.InvitedSlot, error) {
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
			"invitedSlots": bson.M{
				"slotId": slotId,
			},
		},
	}

	// Perform the update operation
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
