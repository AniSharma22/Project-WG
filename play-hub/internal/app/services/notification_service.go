package services

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	repository_interfaces "project2/internal/domain/interfaces/repository"
	service_interfaces "project2/internal/domain/interfaces/service"
	"sync"
)

type NotificationService struct {
	notificationRepo repository_interfaces.NotificationRepository
	notificationWG   *sync.WaitGroup
}

func NewNotificationService(notificationRepo repository_interfaces.NotificationRepository) service_interfaces.NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
		notificationWG:   &sync.WaitGroup{},
	}
}

func (n *NotificationService) GetUserNotifications(ctx context.Context, userId uuid.UUID) ([]entities.Notification, error) {
	return n.notificationRepo.FetchUserNotifications(ctx, userId)
}
