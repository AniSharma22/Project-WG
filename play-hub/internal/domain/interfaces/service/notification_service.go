package service_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
)

type NotificationService interface {
	GetUserNotifications(ctx context.Context, userId uuid.UUID) ([]entities.Notification, error)
}
