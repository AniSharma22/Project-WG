package services

import (
	"project2/internal/domain/interfaces"
)

type NotificationService struct {
	notificationRepo interfaces.NotificationRepository
}

func NewNotificationService(notificationRepo interfaces.NotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
	}
}
