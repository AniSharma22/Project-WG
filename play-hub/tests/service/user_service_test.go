package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/internal/models"
	"project2/pkg/globals"
	"testing"
	"time"
)

func TestUserService_Signup(t *testing.T) {
	tests := []struct {
		name          string
		newUser       *entities.User
		mockRepoError error
		expectedRole  string
		expectedError bool
	}{
		{
			name: "Successful Signup",
			newUser: &entities.User{
				Email:    "test@example.com",
				Password: "TestPassword",
				PhoneNo:  "8989898989",
				Gender:   "male",
			},
			mockRepoError: nil,
			expectedRole:  "user",
			expectedError: false,
		},
		{
			name: "Signup Failure",
			newUser: &entities.User{
				Email:    "test@example.com",
				Password: "TestPassword",
				PhoneNo:  "8989898989",
				Gender:   "male",
			},
			mockRepoError: errors.New("mock repository error"),
			expectedRole:  "user",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().CreateUser(gomock.AssignableToTypeOf(tt.newUser)).Return(tt.mockRepoError).Times(1)

			err := userService.Signup(tt.newUser)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedRole, tt.newUser.Role)
			assert.Empty(t, tt.newUser.InvitedSlots)
		})
	}
}

func TestUserService_EmailAlreadyExists(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		mockRepoError error
		expected      bool
	}{
		{
			name:          "Email Exists",
			email:         "test@example.com",
			mockRepoError: nil,
			expected:      true,
		},
		{
			name:          "Email Does Not Exist",
			email:         "test@example.com",
			mockRepoError: errors.New("email does not exist"),
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().EmailAlreadyExists(tt.email).Return(tt.mockRepoError).Times(1)

			exists := userService.EmailAlreadyExists(tt.email)

			assert.Equal(t, tt.expected, exists)
		})
	}
}

func TestUserService_Login(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		password      []byte
		mockUser      *entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name:     "Successful Login",
			email:    "test@example.com",
			password: []byte("TestPassword"),
			mockUser: &entities.User{
				Email:    "test@example.com",
				Password: "$2a$14$UrwH7llst0N6/k0NkdHktub1zWCosEYzFpkIX8.mJF87BNEZ3LkUu", // Assumes the password check is internal
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:     "Wrong Password",
			email:    "test@example.com",
			password: []byte("WrongPassword"),
			mockUser: &entities.User{
				Email:    "test@example.com",
				Password: "hashedPassword",
			},
			mockRepoError: nil,
			expectedError: true,
		},
		{
			name:          "User Not Found",
			email:         "test@example.com",
			password:      []byte("TestPassword"),
			mockUser:      nil,
			mockRepoError: errors.New("user not found"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().GetUserByEmail(tt.email).Return(tt.mockUser, tt.mockRepoError).Times(1)

			user, err := userService.Login(tt.email, tt.password)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}
		})
	}
}

func TestUserService_GetUserById(t *testing.T) {
	tests := []struct {
		name          string
		userId        primitive.ObjectID
		mockUser      *entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name:   "User Found",
			userId: primitive.NewObjectID(),
			mockUser: &entities.User{
				ID:    primitive.NewObjectID(),
				Email: "test@example.com",
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "User Not Found",
			userId:        primitive.NewObjectID(),
			mockUser:      nil,
			mockRepoError: errors.New("user not found"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().GetUserById(tt.userId).Return(tt.mockUser, tt.mockRepoError).Times(1)

			user, err := userService.GetUserById(tt.userId)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockUser, user)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		mockUser      *entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name:  "User Found",
			email: "test@example.com",
			mockUser: &entities.User{
				Email:    "test@example.com",
				Password: "hashedPassword",
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "User Not Found",
			email:         "test@example.com",
			mockUser:      nil,
			mockRepoError: errors.New("user not found"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().GetUserByEmail(tt.email).Return(tt.mockUser, tt.mockRepoError).Times(1)

			user, err := userService.GetUserByEmail(tt.email)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockUser, user)
			}
		})
	}
}

func TestUserService_GetPendingInvites(t *testing.T) {
	slotId := primitive.NewObjectID()
	gameId := primitive.NewObjectID()
	date := time.Now()
	startTime := time.Now().Add(2 * time.Hour)
	endTime := time.Now().Add(2*time.Hour + 20*time.Minute)

	tests := []struct {
		name            string
		mockInvites     []primitive.ObjectID
		mockSlot        *entities.Slot
		mockGame        *entities.Game
		mockRepoError   error
		expectedError   bool
		expectedInvites []models.Invite
	}{
		{
			name: "Successful Retrieval",
			mockInvites: []primitive.ObjectID{
				slotId,
			},
			mockSlot: &entities.Slot{
				ID:        slotId,
				GameID:    gameId,
				Date:      date,
				StartTime: startTime,
				EndTime:   endTime,
			},
			mockGame: &entities.Game{
				ID:   gameId,
				Name: "Test Game",
			},
			mockRepoError: nil,
			expectedError: false,
			expectedInvites: []models.Invite{
				{
					SlotId:      slotId,
					GameName:    "Test Game",
					Date:        date,
					StartTime:   startTime,
					EndTime:     endTime,
					BookedUsers: []string{},
				},
			},
		},
		{
			name:            "Slot Retrieval Error",
			mockInvites:     []primitive.ObjectID{slotId},
			mockSlot:        nil,
			mockGame:        nil,
			mockRepoError:   errors.New("slot not found"),
			expectedError:   true,
			expectedInvites: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Set up the expectation for GetPendingInvites
			mockUserRepo.EXPECT().GetPendingInvites(globals.ActiveUser).Return(tt.mockInvites, tt.mockRepoError).Times(1)

			if tt.mockRepoError == nil {
				mockSlotService.EXPECT().GetSlotById(slotId).Return(tt.mockSlot, nil).Times(1)
				mockGameService.EXPECT().GetGameByID(tt.mockSlot.GameID).Return(tt.mockGame, nil).Times(1)
			}

			// Call the method under test
			invites, err := userService.GetPendingInvites()

			// Assert results based on expectations
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, invites)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedInvites[0].SlotId, invites[0].SlotId)
				assert.Equal(t, tt.expectedInvites[0].GameName, invites[0].GameName)
			}
		})
	}
}

func TestUserService_AcceptInvite(t *testing.T) {
	slotId := primitive.NewObjectID()

	tests := []struct {
		name                   string
		mockSlotRetrievalError error
		mockSlot               *entities.Slot
		mockGameRetrievalError error
		mockGame               *entities.Game
		mockBookSlotError      error
		mockUserRepoError      error
		expectedError          bool
	}{
		{
			name:                   "Successful Acceptance",
			mockSlotRetrievalError: nil,
			mockSlot: &entities.Slot{
				ID:     slotId,
				GameID: primitive.NewObjectID(),
			},
			mockGameRetrievalError: nil,
			mockGame: &entities.Game{
				ID:   primitive.NewObjectID(),
				Name: "Test Game",
			},
			mockBookSlotError: nil,
			mockUserRepoError: nil,
			expectedError:     false,
		},
		{
			name:                   "Slot Retrieval Error",
			mockSlotRetrievalError: errors.New("slot not found"),
			mockSlot:               nil,
			mockGameRetrievalError: nil,
			mockGame:               nil,
			mockBookSlotError:      nil,
			mockUserRepoError:      nil,
			expectedError:          true,
		},
		{
			name:                   "Game Retrieval Error",
			mockSlotRetrievalError: nil,
			mockSlot: &entities.Slot{
				ID:     slotId,
				GameID: primitive.NewObjectID(),
			},
			mockGameRetrievalError: errors.New("game not found"),
			mockGame:               nil,
			mockBookSlotError:      nil,
			mockUserRepoError:      nil,
			expectedError:          true,
		},
		{
			name:                   "Booking Slot Error",
			mockSlotRetrievalError: nil,
			mockSlot: &entities.Slot{
				ID:     slotId,
				GameID: primitive.NewObjectID(),
			},
			mockGameRetrievalError: nil,
			mockGame: &entities.Game{
				ID:   primitive.NewObjectID(),
				Name: "Test Game",
			},
			mockBookSlotError: errors.New("slot booking error"),
			mockUserRepoError: nil,
			expectedError:     true,
		},
		{
			name:                   "Delete Invite Error",
			mockSlotRetrievalError: nil,
			mockSlot: &entities.Slot{
				ID:     slotId,
				GameID: primitive.NewObjectID(),
			},
			mockGameRetrievalError: nil,
			mockGame: &entities.Game{
				ID:   primitive.NewObjectID(),
				Name: "Test Game",
			},
			mockBookSlotError: nil,
			mockUserRepoError: errors.New("delete invite error"),
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Mock Slot Retrieval
			mockSlotService.EXPECT().
				GetSlotById(slotId).
				Return(tt.mockSlot, tt.mockSlotRetrievalError).
				Times(1)

			// Only mock Game Retrieval if Slot Retrieval is successful
			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil {
				mockGameService.EXPECT().
					GetGameByID(tt.mockSlot.GameID).
					Return(tt.mockGame, tt.mockGameRetrievalError).
					Times(1)
			}

			// Only mock BookSlot if both Slot and Game Retrievals are successful
			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil && tt.mockGameRetrievalError == nil && tt.mockGame != nil {
				mockSlotService.EXPECT().
					BookSlot(tt.mockGame, tt.mockSlot).
					Return(tt.mockBookSlotError).
					Times(1)
			}

			// Only mock DeleteInvite if Slot booking is successful
			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil &&
				tt.mockGameRetrievalError == nil && tt.mockGame != nil &&
				tt.mockBookSlotError == nil {
				mockUserRepo.EXPECT().
					DeleteInvite(slotId).
					Return(tt.mockUserRepoError).
					Times(1)
			}

			err := userService.AcceptInvite(slotId)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_RejectInvite(t *testing.T) {
	slotId := primitive.NewObjectID()

	tests := []struct {
		name          string
		mockRepoError error
		expectedError bool
	}{
		{
			name:          "Successful Rejection",
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "Repository Error",
			mockRepoError: errors.New("error deleting invite"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().DeleteInvite(slotId).Return(tt.mockRepoError).Times(1)

			err := userService.RejectInvite(slotId)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_AddResult(t *testing.T) {
	userId := primitive.NewObjectID()

	tests := []struct {
		name          string
		result        string
		mockRepoError error
		expectedError bool
	}{
		{
			name:          "Add Win",
			result:        "win",
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "Add Loss",
			result:        "loss",
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "Invalid Result",
			result:        "invalid",
			mockRepoError: nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			if tt.result == "win" {
				mockUserRepo.EXPECT().AddWin(userId).Return(tt.mockRepoError).Times(1)
			} else if tt.result == "loss" {
				mockUserRepo.EXPECT().AddLoss(userId).Return(tt.mockRepoError).Times(1)
			}

			err := userService.AddResult(userId, tt.result)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_GetAllUsersByScore(t *testing.T) {
	tests := []struct {
		name          string
		mockUsers     []entities.User
		mockRepoError error
		expectedError bool
		expectedUsers []entities.User
	}{
		{
			name: "Successful Retrieval",
			mockUsers: []entities.User{
				{Email: "test1@example.com"},
				{Email: "test2@example.com"},
			},
			mockRepoError: nil,
			expectedError: false,
			expectedUsers: []entities.User{
				{Email: "test1@example.com"},
				{Email: "test2@example.com"},
			},
		},
		{
			name:          "Repository Error",
			mockUsers:     nil,
			mockRepoError: errors.New("repository error"),
			expectedError: true,
			expectedUsers: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().GetAllUsersByScore().Return(tt.mockUsers, tt.mockRepoError).Times(1)

			users, err := userService.GetAllUsersByScore()

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedUsers, users)
			}
		})
	}
}
