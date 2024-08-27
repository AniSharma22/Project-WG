package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/models"
)

type MockUserService struct {
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (m *MockUserService) Signup(user *entities.User) error {
	return nil
}

func (m *MockUserService) EmailAlreadyExists(email string) bool {
	return false
}

func (m *MockUserService) Login(email string, password []byte) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserService) GetUserById(userId primitive.ObjectID) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserService) GetUserByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserService) GetPendingInvites() ([]models.Invite, error) {
	return nil, nil
}

func (m *MockUserService) AcceptInvite(slotId primitive.ObjectID) error {
	return nil
}

func (m *MockUserService) RejectInvite(slotId primitive.ObjectID) error {
	return nil
}

func (m *MockUserService) AddResult(userId primitive.ObjectID, result string) error {
	return nil
}

func (m *MockUserService) GetAllUsersByScore() ([]entities.User, error) {
	return nil, nil
}
