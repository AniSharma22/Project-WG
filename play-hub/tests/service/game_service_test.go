package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"testing"
)

func TestGameService_GetGameByID(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameID := primitive.NewObjectID()
	mockGame := &entities.Game{ID: gameID, Name: "Test Game", MaxCapacity: 4}

	tests := []struct {
		name          string
		mockGame      *entities.Game
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Retrieval",
			mockGame:      mockGame,
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Game Not Found",
			mockGame:      nil,
			mockError:     errors.New("game not found"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGameRepo.EXPECT().GetGameByID(gameID).Return(tt.mockGame, tt.mockError).Times(1)

			game, err := gameService.GetGameByID(gameID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockGame, game)
			}
		})
	}
}

func TestGameService_GetAllGames(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockGames := []entities.Game{
		{ID: primitive.NewObjectID(), Name: "Game 1", MaxCapacity: 4},
		{ID: primitive.NewObjectID(), Name: "Game 2", MaxCapacity: 2},
	}

	tests := []struct {
		name          string
		mockGames     []entities.Game
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Retrieval",
			mockGames:     mockGames,
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error in Retrieval",
			mockGames:     nil,
			mockError:     errors.New("failed to retrieve games"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGameRepo.EXPECT().GetAllGames().Return(tt.mockGames, tt.mockError).Times(1)

			games, err := gameService.GetAllGames()

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockGames, games)
			}
		})
	}
}

func TestGameService_CreateGame(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameName := "New Game"
	maxPlayers := 4

	tests := []struct {
		name          string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Creation",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error in Creation",
			mockError:     errors.New("failed to create game"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGameRepo.EXPECT().CreateGame(gomock.Any()).Return(tt.mockError).Times(1)

			err := gameService.CreateGame(gameName, maxPlayers)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGameService_DeleteGame(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameID := primitive.NewObjectID()

	tests := []struct {
		name          string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Deletion",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error in Deletion",
			mockError:     errors.New("failed to delete game"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGameRepo.EXPECT().DeleteGame(gameID).Return(tt.mockError).Times(1)

			err := gameService.DeleteGame(gameID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
