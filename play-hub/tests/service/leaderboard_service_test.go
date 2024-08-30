package service_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"testing"
)

func TestLeaderboardService_GetOverallLeaderboard(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUsers := []entities.User{
		{ID: primitive.NewObjectID(), Email: "User 1", OverallScore: 100},
		{ID: primitive.NewObjectID(), Email: "User 2", OverallScore: 90},
	}

	tests := []struct {
		name          string
		mockUsers     []entities.User
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Retrieval",
			mockUsers:     mockUsers,
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error in Retrieval",
			mockUsers:     nil,
			mockError:     errors.New("failed to retrieve leaderboard"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService.EXPECT().GetAllUsersByScore().Return(tt.mockUsers, tt.mockError).Times(1)

			users, err := leaderboardService.GetOverallLeaderboard()

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockUsers, users)
			}
		})
	}
}
