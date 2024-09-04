package service_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain/entities"
	"testing"
)

func TestGameService_GetGameByID(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameID := uuid.New()
	expectedGame := &entities.Game{GameID: gameID}

	mockGameRepo.EXPECT().
		FetchGameByID(gomock.Any(), gameID).
		Return(expectedGame, nil).
		Times(1)

	game, err := gameService.GetGameByID(context.Background(), gameID)

	assert.NoError(t, err)
	assert.Equal(t, expectedGame, game)
}

func TestGameService_GetAllGames(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	expectedGames := []entities.Game{
		{GameID: uuid.New()},
		{GameID: uuid.New()},
	}

	mockGameRepo.EXPECT().
		FetchAllGames(gomock.Any()).
		Return(expectedGames, nil).
		Times(1)

	games, err := gameService.GetAllGames(context.Background())

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedGames, games)
}

func TestGameService_DeleteGame(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameID := uuid.New()

	mockGameRepo.EXPECT().
		DeleteGame(gomock.Any(), gameID).
		Return(nil).
		Times(1)

	err := gameService.DeleteGame(context.Background(), gameID)

	assert.NoError(t, err)
}

func TestGameService_CreateGame(t *testing.T) {
	mockID := uuid.New()
	tests := []struct {
		name          string
		game          *entities.Game
		mockSetup     func()
		expectedID    uuid.UUID
		expectedError error
	}{
		{
			name: "Success",
			game: &entities.Game{},
			mockSetup: func() {
				mockGameRepo.EXPECT().
					CreateGame(gomock.Any(), gomock.Any()).
					Return(mockID, nil).
					Times(1)
			},
			expectedID:    mockID,
			expectedError: nil,
		},
		{
			name: "Error",
			game: &entities.Game{},
			mockSetup: func() {
				mockGameRepo.EXPECT().
					CreateGame(gomock.Any(), gomock.Any()).
					Return(uuid.Nil, errors.New("creation error")).
					Times(1)
			},
			expectedID:    uuid.Nil,
			expectedError: fmt.Errorf("failed to create game: creation error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			tt.mockSetup()

			id, _ := gameService.CreateGame(context.Background(), tt.game)

			assert.Equal(t, tt.expectedID, id)
			//assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestGameService_UpdateGameStatus(t *testing.T) {
	tests := []struct {
		name          string
		gameID        uuid.UUID
		status        bool
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "Success",
			gameID: uuid.New(),
			status: true,
			mockSetup: func() {
				mockGameRepo.EXPECT().
					FetchGameByID(gomock.Any(), gomock.Any()).
					Return(&entities.Game{GameID: uuid.New()}, nil).
					Times(1)
				mockGameRepo.EXPECT().
					UpdateGameStatus(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
			},
			expectedError: nil,
		},
		{
			name:   "GameNotFound",
			gameID: uuid.New(),
			status: true,
			mockSetup: func() {
				mockGameRepo.EXPECT().
					FetchGameByID(gomock.Any(), gomock.Any()).
					Return(nil, nil).
					Times(1)
			},
			expectedError: errors.New("game not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			tt.mockSetup()

			err := gameService.UpdateGameStatus(context.Background(), tt.gameID, tt.status)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}

//func TestGameService_GetGameByID(t *testing.T) {
//	teardown := setup(t)
//	defer teardown()
//
//	gameID := primitive.NewObjectID()
//	mockGame := &entities.Game{ID: gameID, Name: "Test Game", MaxCapacity: 4}
//
//	tests := []struct {
//		name          string
//		mockGame      *entities.Game
//		mockError     error
//		expectedError bool
//	}{
//		{
//			name:          "Successful Retrieval",
//			mockGame:      mockGame,
//			mockError:     nil,
//			expectedError: false,
//		},
//		{
//			name:          "Game Not Found",
//			mockGame:      nil,
//			mockError:     errors.New("game not found"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockGameRepo.EXPECT().GetGameByID(gameID).Return(tt.mockGame, tt.mockError).Times(1)
//
//			game, err := gameService.GetGameByID(gameID)
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.mockGame, game)
//			}
//		})
//	}
//}
//
//func TestGameService_GetAllGames(t *testing.T) {
//	teardown := setup(t)
//	defer teardown()
//
//	mockGames := []entities.Game{
//		{ID: primitive.NewObjectID(), Name: "Game 1", MaxCapacity: 4},
//		{ID: primitive.NewObjectID(), Name: "Game 2", MaxCapacity: 2},
//	}
//
//	tests := []struct {
//		name          string
//		mockGames     []entities.Game
//		mockError     error
//		expectedError bool
//	}{
//		{
//			name:          "Successful Retrieval",
//			mockGames:     mockGames,
//			mockError:     nil,
//			expectedError: false,
//		},
//		{
//			name:          "Error in Retrieval",
//			mockGames:     nil,
//			mockError:     errors.New("failed to retrieve games"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockGameRepo.EXPECT().GetAllGames().Return(tt.mockGames, tt.mockError).Times(1)
//
//			games, err := gameService.GetAllGames()
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.mockGames, games)
//			}
//		})
//	}
//}
//
//func TestGameService_CreateGame(t *testing.T) {
//	teardown := setup(t)
//	defer teardown()
//
//	gameName := "New Game"
//	maxPlayers := 4
//
//	tests := []struct {
//		name          string
//		mockError     error
//		expectedError bool
//	}{
//		{
//			name:          "Successful Creation",
//			mockError:     nil,
//			expectedError: false,
//		},
//		{
//			name:          "Error in Creation",
//			mockError:     errors.New("failed to create game"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockGameRepo.EXPECT().CreateGame(gomock.Any()).Return(tt.mockError).Times(1)
//
//			err := gameService.CreateGame(gameName, maxPlayers)
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestGameService_DeleteGame(t *testing.T) {
//	teardown := setup(t)
//	defer teardown()
//
//	gameID := primitive.NewObjectID()
//
//	tests := []struct {
//		name          string
//		mockError     error
//		expectedError bool
//	}{
//		{
//			name:          "Successful Deletion",
//			mockError:     nil,
//			expectedError: false,
//		},
//		{
//			name:          "Error in Deletion",
//			mockError:     errors.New("failed to delete game"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockGameRepo.EXPECT().DeleteGame(gameID).Return(tt.mockError).Times(1)
//
//			err := gameService.DeleteGame(gameID)
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
