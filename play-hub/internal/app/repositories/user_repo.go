package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	interfaces "project2/internal/domain/interfaces/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepo{
		db: db,
	}
}

// CreateUser creates a new user in the DB
func (r *userRepo) CreateUser(ctx context.Context, user *entities.User) (uuid.UUID, error) {
	// Insert into PostgresSQL and return the user_id
	query := `
		INSERT INTO users (username, email, password, mobile_number, gender)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id
	`
	row := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.MobileNumber, user.Gender)

	// Variable to hold the returned user_id
	var userID uuid.UUID
	err := row.Scan(&userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert user into PostgreSQL and retrieve user_id: %w", err)
	}

	return userID, nil
}

// FetchUserByEmail retrieves a user by their email address.
func (r *userRepo) FetchUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT user_id, username, email, password, mobile_number, gender,role FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var user entities.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.MobileNumber, &user.Gender, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
	}

	return &user, nil
}

// FetchUserById retrieves a user by their unique user_id.
func (r *userRepo) FetchUserById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `SELECT user_id, username, email, password, mobile_number, gender,role FROM users WHERE user_id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var user entities.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.MobileNumber, &user.Gender, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("failed to fetch user by ID: %w", err)
	}

	return &user, nil
}

// FetchAllUsers retrieves all users from the database.
func (r *userRepo) FetchAllUsers(ctx context.Context) ([]entities.User, error) {
	query := `SELECT user_id, username, email, password, mobile_number, gender FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all users: %w", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.MobileNumber, &user.Gender); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %w", err)
	}

	return users, nil
}

// EmailAlreadyExists checks if the given email already exists in the database.
func (r *userRepo) EmailAlreadyExists(ctx context.Context, email string) bool {
	query := `SELECT 1 FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var exists bool
	err := row.Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return exists
}

func (r *userRepo) FetchUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := `SELECT user_id, username, email, password, mobile_number, gender, role, created_at, updated_at FROM users WHERE username = $1`
	row := r.db.QueryRowContext(ctx, query, username)

	var user entities.User
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.MobileNumber,
		&user.Gender,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("failed to fetch user by username: %w", err)
	}

	return &user, nil
}

//func (r *userRepo) CreateUser(user *entities.User) error {
//	_, err := r.collection.InsertOne(context.Background(), user)
//	if err != nil {
//		fmt.Println("Error inserting user:", err)
//		return err
//	}
//	return nil
//}

//func (r *userRepo) AddToInvites(userId primitive.ObjectID, slotId primitive.ObjectID) error {
//	// Create a filter to find the user by ID
//	filter := bson.M{"_id": userId}
//
//	// Create an update to add the invite to the InvitedSlots array
//	update := bson.M{
//		"$push": bson.M{
//			"invitedSlots": slotId,
//		},
//	}
//
//	// Perform the update operation
//	_, err := r.collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		fmt.Println("Error updating user with invited slot:", err)
//		return err
//	}
//
//	return nil
//}
//
//func (r *userRepo) GetUserByEmail(email string) (*entities.User, error) {
//	filter := bson.M{"email": email}
//
//	var user entities.User
//	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			return nil, fmt.Errorf("user not found for email: %s", email)
//		}
//		return nil, fmt.Errorf("error querying database: %w", err)
//	}
//
//	if user.ID.IsZero() {
//		return nil, fmt.Errorf("user found but ID is zero for email: %s", email)
//	}
//
//	return &user, nil
//}
//
//// EmailAlreadyExists checks if the email already exists in the database.
//func (r *userRepo) EmailAlreadyExists(email string) error {
//	// Create a filter to find a document with the specified email
//	filter := bson.M{"email": email}
//
//	// Perform the query to find one document with the specified email
//	var result entities.User
//	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
//
//	if errors.Is(err, mongo.ErrNoDocuments) {
//		// No document found, email does not exist
//		return nil
//	} else if err != nil {
//		// Error occurred while querying
//		return err
//	}
//
//	fmt.Println(result)
//
//	// Email already exists
//	return nil
//}
//
//func (r *userRepo) AddWin(userId primitive.ObjectID) error {
//	// Retrieve the current user stats
//	var user entities.User
//	err := r.collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&user)
//	if err != nil {
//		fmt.Println("Error fetching user data:", err)
//		return err
//	}
//
//	// Calculate new score
//	newScore := utils.GetTotalScore(user.Wins+1, user.Losses)
//
//	// Update the user's stats
//	filter := bson.D{{Key: "_id", Value: userId}}
//	update := bson.D{
//		{Key: "$inc", Value: bson.D{
//			{Key: "wins", Value: 1},
//			{Key: "overallScore", Value: newScore - user.OverallScore},
//		}},
//	}
//
//	_, err = r.collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		fmt.Println("Error updating user wins:", err)
//		return err
//	}
//	return nil
//}
//
//func (r *userRepo) AddLoss(userId primitive.ObjectID) error {
//	// Retrieve the current user stats
//	var user entities.User
//	err := r.collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&user)
//	if err != nil {
//		fmt.Println("Error fetching user data:", err)
//		return err
//	}
//
//	// Calculate new score
//	newScore := utils.GetTotalScore(user.Wins, user.Losses+1)
//
//	// Update the user's stats
//	filter := bson.D{{Key: "_id", Value: userId}}
//	update := bson.D{
//		{Key: "$inc", Value: bson.D{
//			{Key: "losses", Value: 1},
//			{Key: "overallScore", Value: newScore - user.OverallScore}, // Adjust overallScore based on new calculation
//		}},
//	}
//
//	_, err = r.collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		fmt.Println("Error updating user losses:", err)
//		return err
//	}
//	return nil
//}
//
//func (r *userRepo) GetAllUsers() ([]entities.User, error) {
//	cursor, err := r.collection.Find(context.Background(), bson.D{})
//	if err != nil {
//		fmt.Println("Error finding users:", err)
//		return nil, err
//	}
//	defer cursor.Close(context.Background())
//
//	var users []entities.User
//	for cursor.Next(context.Background()) {
//		var user entities.User
//		if err := cursor.Decode(&user); err != nil {
//			fmt.Println("Error decoding user:", err)
//			return nil, err
//		}
//		users = append(users, user)
//	}
//	if err := cursor.Err(); err != nil {
//		fmt.Println("Cursor error:", err)
//		return nil, err
//	}
//
//	return users, nil
//}
//
//func (r *userRepo) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
//	filter := bson.D{{"_id", userId}}
//	var user entities.User
//	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			return nil, fmt.Errorf("user not found for id: %s", userId)
//
//		}
//	}
//	return &user, nil
//}
//
//func (r *userRepo) GetPendingInvites(email string) ([]primitive.ObjectID, error) {
//	filter := bson.M{"email": email}
//	var user entities.User
//
//	// Find the user document
//	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			return nil, fmt.Errorf("user not found for email: %s", email)
//		}
//		return nil, fmt.Errorf("error fetching user: %w", err)
//	}
//
//	// Return the invites
//	return user.InvitedSlots, nil
//}
//
//func (r *userRepo) DeleteInvite(slotId primitive.ObjectID) error {
//	email := globals.ActiveUser
//
//	// Define the filter to find the user by email
//	filter := bson.M{"email": email}
//
//	// Define the update to pull the slotId from the invitedSlots array
//	update := bson.M{
//		"$pull": bson.M{
//			"invitedSlots": slotId,
//		},
//	}
//
//	// Perform the update operation
//	_, err := r.collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *userRepo) GetAllUsersByScore() ([]entities.User, error) {
//	var users []entities.User
//
//	// Define the filter to only include users with the role "user"
//	filter := bson.M{"role": "user"}
//
//	// Define the sort filter to sort by OverallScore in descending order
//	opts := options.Find().SetSort(bson.M{"overallScore": -1})
//
//	// Perform the find operation with the filter and sort options
//	cursor, err := r.collection.Find(context.Background(), filter, opts)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(context.Background())
//
//	// Iterate through the cursor and decode each user
//	for cursor.Next(context.Background()) {
//		var user entities.User
//		if err := cursor.Decode(&user); err != nil {
//			return nil, err
//		}
//		users = append(users, user)
//	}
//
//	// Check for errors during cursor iteration
//	if err := cursor.Err(); err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}
