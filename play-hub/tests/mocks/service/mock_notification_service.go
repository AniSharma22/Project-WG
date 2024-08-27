package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
)

type MockNotificationService struct {
}

func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{}
}

func (n *MockNotificationService) GetUserNotifications(userId primitive.ObjectID) ([]entities.Notification, error) {
	return nil, nil
}
