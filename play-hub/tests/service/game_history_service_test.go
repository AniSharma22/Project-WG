package service_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/pkg/globals"
	"testing"
	"time"
)

func TestGameHistoryService_GetTotalGameHistory(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUser := &entities.User{ID: primitive.NewObjectID(), Email: globals.ActiveUser}
	mockGameHistories := []entities.GameHistory{
		{ID: primitive.NewObjectID(), SlotID: primitive.NewObjectID()},
	}

	tests := []struct {
		name              string
		mockUserError     error
		mockHistoryError  error
		expectedError     bool
		expectedHistories []entities.GameHistory
	}{
		{
			name:              "Successful Retrieval",
			mockUserError:     nil,
			mockHistoryError:  nil,
			expectedError:     false,
			expectedHistories: mockGameHistories,
		},
		{
			name:              "User Retrieval Error",
			mockUserError:     errors.New("user not found"),
			mockHistoryError:  nil,
			expectedError:     true,
			expectedHistories: nil,
		},
		{
			name:              "Game History Retrieval Error",
			mockUserError:     nil,
			mockHistoryError:  errors.New("history not found"),
			expectedError:     true,
			expectedHistories: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService.EXPECT().GetUserByEmail(globals.ActiveUser).Return(mockUser, tt.mockUserError).Times(1)
			if tt.mockUserError == nil {
				mockGameHistoryRepo.EXPECT().GetUserGameHistory(mockUser.ID).Return(tt.expectedHistories, tt.mockHistoryError).Times(1)
			}

			histories, err := gameHistoryService.GetTotalGameHistory()

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, histories)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedHistories, histories)
			}
		})
	}
}

func TestGameHistoryService_GetResultsToUpdate(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUser := &entities.User{ID: primitive.NewObjectID(), Email: globals.ActiveUser}
	mockGameHistories := []entities.GameHistory{
		{ID: primitive.NewObjectID(), SlotID: primitive.NewObjectID()},
	}
	mockSlot := &entities.Slot{
		ID:      primitive.NewObjectID(),
		EndTime: time.Now().Add(-time.Hour), // End time in the past
	}

	tests := []struct {
		name              string
		mockUserError     error
		mockHistoryError  error
		mockSlotError     error
		expectedError     bool
		expectedHistories []entities.GameHistory
	}{
		{
			name:              "Successful Update Retrieval",
			mockUserError:     nil,
			mockHistoryError:  nil,
			mockSlotError:     nil,
			expectedError:     false,
			expectedHistories: mockGameHistories,
		},
		{
			name:              "User Retrieval Error",
			mockUserError:     errors.New("user not found"),
			mockHistoryError:  nil,
			mockSlotError:     nil,
			expectedError:     true,
			expectedHistories: nil,
		},
		{
			name:              "Game History Retrieval Error",
			mockUserError:     nil,
			mockHistoryError:  errors.New("history not found"),
			mockSlotError:     nil,
			expectedError:     true,
			expectedHistories: nil,
		},
		{
			name:              "Slot Retrieval Error",
			mockUserError:     nil,
			mockHistoryError:  nil,
			mockSlotError:     errors.New("slot not found"),
			expectedError:     false, // The function should continue without returning an error
			expectedHistories: nil,   // No results to update due to slot retrieval failure
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService.EXPECT().GetUserByEmail(globals.ActiveUser).Return(mockUser, tt.mockUserError).Times(1)
			if tt.mockUserError == nil {
				mockGameHistoryRepo.EXPECT().GetResultsToUpdate(mockUser.ID).Return(mockGameHistories, tt.mockHistoryError).Times(1)
				if tt.mockHistoryError == nil {
					mockSlotService.EXPECT().GetSlotById(mockGameHistories[0].SlotID).Return(mockSlot, tt.mockSlotError).Times(1)
				}
			}

			resultsToUpdate, err := gameHistoryService.GetResultsToUpdate()

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resultsToUpdate)
			} else {
				assert.NoError(t, err)
				if tt.mockSlotError == nil {
					assert.Equal(t, tt.expectedHistories, resultsToUpdate)
				} else {
					assert.Empty(t, resultsToUpdate) // No results to update due to slot retrieval error
				}
			}
		})
	}
}

func TestGameHistoryService_UpdateResult(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUser := &entities.User{ID: primitive.NewObjectID(), Email: globals.ActiveUser}
	slotId := primitive.NewObjectID()
	result := "win"

	tests := []struct {
		name                string
		mockUserError       error
		mockUpdateError     error
		mockAddResultError  error
		mockSlotUpdateError error
		expectedError       bool
	}{
		{
			name:                "Successful Update",
			mockUserError:       nil,
			mockUpdateError:     nil,
			mockAddResultError:  nil,
			mockSlotUpdateError: nil,
			expectedError:       false,
		},
		{
			name:                "User Retrieval Error",
			mockUserError:       errors.New("user not found"),
			mockUpdateError:     nil,
			mockAddResultError:  nil,
			mockSlotUpdateError: nil,
			expectedError:       true,
		},
		{
			name:                "Game History Update Error",
			mockUserError:       nil,
			mockUpdateError:     errors.New("update error"),
			mockAddResultError:  nil,
			mockSlotUpdateError: nil,
			expectedError:       true,
		},
		{
			name:                "Add Result Error",
			mockUserError:       nil,
			mockUpdateError:     nil,
			mockAddResultError:  errors.New("add result error"),
			mockSlotUpdateError: nil,
			expectedError:       true,
		},
		{
			name:                "Slot Update Error",
			mockUserError:       nil,
			mockUpdateError:     nil,
			mockAddResultError:  nil,
			mockSlotUpdateError: errors.New("slot update error"),
			expectedError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService.EXPECT().GetUserByEmail(globals.ActiveUser).Return(mockUser, tt.mockUserError).Times(1)
			if tt.mockUserError == nil {
				mockGameHistoryRepo.EXPECT().UpdateResult(result, slotId, mockUser.ID).Return(tt.mockUpdateError).Times(1)
				if tt.mockUpdateError == nil {
					mockUserService.EXPECT().AddResult(mockUser.ID, result).Return(tt.mockAddResultError).Times(1)
					if tt.mockAddResultError == nil {
						mockSlotService.EXPECT().AddResultToSlot(mockUser.ID, slotId, result).Return(tt.mockSlotUpdateError).Times(1)
					}
				}
			}

			err := gameHistoryService.UpdateResult(result, slotId)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
