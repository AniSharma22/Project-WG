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

type notificationRepo struct {
	collection *mongo.Collection
}

func NewNotificationRepo(client *mongo.Client) interfaces.NotificationRepository {
	return &notificationRepo{
		collection: client.Database(config.DB.DBName).Collection(config.DB.NotificationsCollection),
	}
}

// AddNotification adds a new notification to the database.
func (r *notificationRepo) AddNotification(notification *entities.Notification) error {
	_, err := r.collection.InsertOne(context.Background(), notification)
	return err
}

// GetNotifications retrieves notifications for a specific user.
func (r *notificationRepo) GetNotifications(userId primitive.ObjectID) ([]entities.Notification, error) {
	filter := bson.M{"userId": userId}
	opts := options.Find().SetSort(bson.D{{"createdAt", -1}}) // Sort by creation date in descending order

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var notifications []entities.Notification
	if err := cursor.All(context.Background(), &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}

// DeleteNotification deletes a notification by its ID.
func (r *notificationRepo) DeleteNotification(notificationId primitive.ObjectID) error {
	filter := bson.M{"_id": notificationId}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}
