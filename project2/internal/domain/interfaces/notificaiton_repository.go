package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type NotificationRepository interface {
	AddNotification(notification *entities.Notification) error
	GetNotifications(userId primitive.ObjectID) ([]entities.Notification, error)
	DeleteNotification(notificationId primitive.ObjectID) error
}
