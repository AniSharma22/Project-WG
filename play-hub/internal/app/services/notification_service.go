package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
)

type NotificationService struct {
	notificationRepo interfaces.NotificationRepository
}

func NewNotificationService(notificationRepo interfaces.NotificationRepository) interfaces.NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
	}
}

func (n *NotificationService) GetUserNotifications(userId primitive.ObjectID) ([]entities.Notification, error) {
	return n.notificationRepo.GetNotifications(userId)
}
