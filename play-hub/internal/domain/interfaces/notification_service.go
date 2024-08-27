package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type NotificationService interface {
	GetUserNotifications(userId primitive.ObjectID) ([]entities.Notification, error)
}
