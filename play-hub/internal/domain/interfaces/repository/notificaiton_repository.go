package repository_interfaces

import (
	"context"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, notification *entities.Notification) (uuid.UUID, error)
	FetchNotificationByID(ctx context.Context, id uuid.UUID) (*entities.Notification, error)
	FetchUserNotifications(ctx context.Context, userID uuid.UUID) ([]entities.Notification, error)
	MarkNotificationAsRead(ctx context.Context, id uuid.UUID) error
	DeleteNotificationByID(ctx context.Context, id uuid.UUID) error
}
