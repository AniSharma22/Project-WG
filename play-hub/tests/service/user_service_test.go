package service_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain/entities"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"testing"
)

func TestUserService_Signup(t *testing.T) {
	userId, _ := uuid.NewUUID()
	tests := []struct {
		name                       string
		newUser                    *entities.User
		mockCreateUserRepo         error
		mockEmailAlreadyRegistered error
		expectedError              bool
		CreateUserCalled           int
	}{
		{
			name: "Successful Signup",
			newUser: &entities.User{
				Email:        "test.test@watchguard.com",
				Password:     "TestPassword",
				MobileNumber: "8989898989",
				Gender:       "male",
			},
			mockCreateUserRepo:         nil,
			mockEmailAlreadyRegistered: nil,
			expectedError:              false,
			CreateUserCalled:           1,
		},
		{
			name: "Signup Failure",
			newUser: &entities.User{
				Email:        "test.test@watchguard.com",
				Password:     "TestPassword",
				MobileNumber: "8989898989",
				Gender:       "male",
			},
			mockCreateUserRepo:         errors.New("mock repository error"),
			mockEmailAlreadyRegistered: nil,
			expectedError:              true,
			CreateUserCalled:           1,
		},
		{
			name: "Signup Failure - Email Already Registered",
			newUser: &entities.User{
				Email:        "test.test@watchguard.com",
				Password:     "TestPassword",
				MobileNumber: "8989898989",
				Gender:       "male",
			},
			mockCreateUserRepo:         nil,
			mockEmailAlreadyRegistered: errors.New("mock repository error"),
			expectedError:              true,
			CreateUserCalled:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			// Mock the repository call with the expected behavior
			mockUserRepo.EXPECT().
				CreateUser(ctx, tt.newUser).
				Return(userId, tt.mockCreateUserRepo).
				Times(tt.CreateUserCalled)

			// Mock the repository call with the expected behaviour for the emailAlreadyRegistered
			mockUserRepo.EXPECT().
				FetchUserByEmail(ctx, tt.newUser.Email).
				Return(nil, tt.mockEmailAlreadyRegistered).
				Times(1)

			// Call the Signup method
			err := userService.Signup(ctx, tt.newUser)

			// Assert the expected error outcome
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// verify if the global variable is set
			if !tt.expectedError {
				assert.Equal(t, userId, globals.ActiveUser)
			}
		})
	}
}

func TestUserService_EmailAlreadyExists(t *testing.T) {
	tests := []struct {
		name          string
		user          *entities.User
		email         string
		mockRepoError error
		expected      bool
	}{
		{
			name:          "Email Exists",
			user:          &entities.User{},
			email:         "test@example.com",
			mockRepoError: nil,
			expected:      true,
		},
		{
			name:          "error while scanning in repository",
			user:          &entities.User{},
			email:         "test@example.com",
			mockRepoError: errors.New("sql error"),
			expected:      true,
		},
		{
			name:          "Email Does Not Exist",
			user:          nil,
			email:         "test@example.com",
			mockRepoError: nil,
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			mockUserRepo.EXPECT().FetchUserByEmail(ctx, tt.email).Return(tt.user, tt.mockRepoError).Times(1)

			exists := userService.EmailAlreadyRegistered(ctx, tt.email)

			assert.Equal(t, tt.expected, exists)
		})
	}
}

func TestUserService_Login(t *testing.T) {
	userId, _ := uuid.NewUUID()
	hashedPass, _ := utils.GetHashedPassword([]byte("ValidPassword"))

	// Sample user data
	validUser := &entities.User{
		UserID:   userId,
		Email:    "test@example.com",
		Password: hashedPass,
	}

	tests := []struct {
		name               string
		email              string
		password           []byte
		mockFetchUserError error
		mockFetchedUser    *entities.User
		expectedError      bool
		expectedErrorMsg   string
		expectedUserID     uuid.UUID
	}{
		{
			name:               "Successful Login",
			email:              "test@example.com",
			password:           []byte("ValidPassword"),
			mockFetchUserError: nil,
			mockFetchedUser:    validUser,
			expectedError:      false,
			expectedUserID:     userId,
		},
		{
			name:               "User Not Found",
			email:              "notfound@example.com",
			password:           []byte("AnyPassword"),
			mockFetchUserError: nil,
			mockFetchedUser:    nil,
			expectedError:      true,
			expectedErrorMsg:   "user not found",
			expectedUserID:     uuid.Nil,
		},
		{
			name:               "Invalid Password",
			email:              "test@example.com",
			password:           []byte("InvalidPassword"),
			mockFetchUserError: nil,
			mockFetchedUser:    validUser,
			expectedError:      true,
			expectedErrorMsg:   "invalid password",
			expectedUserID:     uuid.Nil,
		},
		{
			name:               "Fetch User Error",
			email:              "test@example.com",
			password:           []byte("ValidPassword"),
			mockFetchUserError: errors.New("database error"),
			mockFetchedUser:    nil,
			expectedError:      true,
			expectedErrorMsg:   "failed to fetch user by email: database error",
			expectedUserID:     uuid.Nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			// Mock the repository call for FetchUserByEmail
			mockUserRepo.EXPECT().
				FetchUserByEmail(ctx, tt.email).
				Return(tt.mockFetchedUser, tt.mockFetchUserError).
				Times(1)

			// Call the Login method
			user, err := userService.Login(ctx, tt.email, tt.password)

			// Assert the expected error outcome
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUserID, user.UserID)
			}

			// Verify if the global variable is set correctly on successful login
			if !tt.expectedError {
				assert.Equal(t, tt.expectedUserID, globals.ActiveUser)
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	userId, _ := uuid.NewUUID()
	// Sample user data
	validUser := &entities.User{
		UserID: userId,
		Email:  "test@example.com",
	}

	tests := []struct {
		name             string
		userID           uuid.UUID
		mockFetchUser    *entities.User
		mockFetchError   error
		expectedUser     *entities.User
		expectedErrorMsg string
		expectedError    bool
	}{
		{
			name:           "User Found",
			userID:         userId,
			mockFetchUser:  validUser,
			mockFetchError: nil,
			expectedUser:   validUser,
			expectedError:  false,
		},
		{
			name:             "User Not Found",
			userID:           uuid.New(),
			mockFetchUser:    nil,
			mockFetchError:   errors.New("user not found"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "user not found",
		},
		{
			name:             "Repository Error",
			userID:           uuid.New(),
			mockFetchUser:    nil,
			mockFetchError:   errors.New("database error"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			// Mock the repository call for FetchUserById
			mockUserRepo.EXPECT().
				FetchUserById(ctx, tt.userID).
				Return(tt.mockFetchUser, tt.mockFetchError).
				Times(1)

			// Call the GetUserByID method
			user, err := userService.GetUserByID(ctx, tt.userID)

			// Assert the expected outcome
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	userId, _ := uuid.NewUUID()
	// Sample user data
	validUser := &entities.User{
		UserID: userId,
		Email:  "test@example.com",
	}

	tests := []struct {
		name             string
		email            string
		mockFetchUser    *entities.User
		mockFetchError   error
		expectedUser     *entities.User
		expectedErrorMsg string
		expectedError    bool
	}{
		{
			name:           "User Found",
			email:          "test@example.com",
			mockFetchUser:  validUser,
			mockFetchError: nil,
			expectedUser:   validUser,
			expectedError:  false,
		},
		{
			name:             "User Not Found",
			email:            "notfound@example.com",
			mockFetchUser:    nil,
			mockFetchError:   errors.New("user not found"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "user not found",
		},
		{
			name:             "Repository Error",
			email:            "test@example.com",
			mockFetchUser:    nil,
			mockFetchError:   errors.New("database error"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			// Mock the repository call for FetchUserByEmail
			mockUserRepo.EXPECT().
				FetchUserByEmail(ctx, tt.email).
				Return(tt.mockFetchUser, tt.mockFetchError).
				Times(1)

			// Call the GetUserByEmail method
			user, err := userService.GetUserByEmail(ctx, tt.email)

			// Assert the expected outcome
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

func TestUserService_GetUserByUsername(t *testing.T) {
	userId, _ := uuid.NewUUID()
	// Sample user data
	validUser := &entities.User{
		UserID:   userId,
		Username: "testuser",
	}

	tests := []struct {
		name             string
		username         string
		mockFetchUser    *entities.User
		mockFetchError   error
		expectedUser     *entities.User
		expectedErrorMsg string
		expectedError    bool
	}{
		{
			name:           "User Found",
			username:       "testuser",
			mockFetchUser:  validUser,
			mockFetchError: nil,
			expectedUser:   validUser,
			expectedError:  false,
		},
		{
			name:             "User Not Found",
			username:         "notfounduser",
			mockFetchUser:    nil,
			mockFetchError:   errors.New("user not found"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "user not found",
		},
		{
			name:             "Repository Error",
			username:         "testuser",
			mockFetchUser:    nil,
			mockFetchError:   errors.New("database error"),
			expectedUser:     nil,
			expectedError:    true,
			expectedErrorMsg: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Prepare the context
			ctx := context.TODO()

			// Mock the repository call for FetchUserByUsername
			mockUserRepo.EXPECT().
				FetchUserByUsername(ctx, tt.username).
				Return(tt.mockFetchUser, tt.mockFetchError).
				Times(1)

			// Call the GetUserByUsername method
			user, err := userService.GetUserByUsername(ctx, tt.username)

			// Assert the expected outcome
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

//func TestUserService_GetUserById(t *testing.T) {
//	tests := []struct {
//		name          string
//		userId        primitive.ObjectID
//		mockUser      *entities.User
//		mockRepoError error
//		expectedError bool
//	}{
//		{
//			name:   "User Found",
//			userId: primitive.NewObjectID(),
//			mockUser: &entities.User{
//				ID:    primitive.NewObjectID(),
//				Email: "test@example.com",
//			},
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "User Not Found",
//			userId:        primitive.NewObjectID(),
//			mockUser:      nil,
//			mockRepoError: errors.New("user not found"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			mockUserRepo.EXPECT().GetUserById(tt.userId).Return(tt.mockUser, tt.mockRepoError).Times(1)
//
//			user, err := userService.GetUserById(tt.userId)
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.mockUser, user)
//			}
//		})
//	}
//}
//
//func TestUserService_GetUserByEmail(t *testing.T) {
//	tests := []struct {
//		name          string
//		email         string
//		mockUser      *entities.User
//		mockRepoError error
//		expectedError bool
//	}{
//		{
//			name:  "User Found",
//			email: "test@example.com",
//			mockUser: &entities.User{
//				Email:    "test@example.com",
//				Password: "hashedPassword",
//			},
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "User Not Found",
//			email:         "test@example.com",
//			mockUser:      nil,
//			mockRepoError: errors.New("user not found"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			mockUserRepo.EXPECT().GetUserByEmail(tt.email).Return(tt.mockUser, tt.mockRepoError).Times(1)
//
//			user, err := userService.GetUserByEmail(tt.email)
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.mockUser, user)
//			}
//		})
//	}
//}
//
//func TestUserService_GetPendingInvites(t *testing.T) {
//	slotId := primitive.NewObjectID()
//	gameId := primitive.NewObjectID()
//	date := time.Now()
//	startTime := time.Now().Add(2 * time.Hour)
//	endTime := time.Now().Add(2*time.Hour + 20*time.Minute)
//
//	tests := []struct {
//		name            string
//		mockInvites     []primitive.ObjectID
//		mockSlot        *entities.Slot
//		mockGame        *entities.Game
//		mockRepoError   error
//		expectedError   bool
//		expectedInvites []models.Invite
//	}{
//		{
//			name: "Successful Retrieval",
//			mockInvites: []primitive.ObjectID{
//				slotId,
//			},
//			mockSlot: &entities.Slot{
//				ID:        slotId,
//				GameID:    gameId,
//				Date:      date,
//				StartTime: startTime,
//				EndTime:   endTime,
//			},
//			mockGame: &entities.Game{
//				ID:   gameId,
//				Name: "Test Game",
//			},
//			mockRepoError: nil,
//			expectedError: false,
//			expectedInvites: []models.Invite{
//				{
//					SlotId:      slotId,
//					GameName:    "Test Game",
//					Date:        date,
//					StartTime:   startTime,
//					EndTime:     endTime,
//					BookedUsers: []string{},
//				},
//			},
//		},
//		{
//			name:            "Slot Retrieval Error",
//			mockInvites:     []primitive.ObjectID{slotId},
//			mockSlot:        nil,
//			mockGame:        nil,
//			mockRepoError:   errors.New("slot not found"),
//			expectedError:   true,
//			expectedInvites: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			// Set up the expectation for GetPendingInvites
//			mockUserRepo.EXPECT().GetPendingInvites(globals.ActiveUser).Return(tt.mockInvites, tt.mockRepoError).Times(1)
//
//			if tt.mockRepoError == nil {
//				mockSlotService.EXPECT().GetSlotById(slotId).Return(tt.mockSlot, nil).Times(1)
//				mockGameService.EXPECT().GetGameByID(tt.mockSlot.GameID).Return(tt.mockGame, nil).Times(1)
//			}
//
//			// Call the method under test
//			invites, err := userService.GetPendingInvites()
//
//			// Assert results based on expectations
//			if tt.expectedError {
//				assert.Error(t, err)
//				assert.Nil(t, invites)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.expectedInvites[0].SlotId, invites[0].SlotId)
//				assert.Equal(t, tt.expectedInvites[0].GameName, invites[0].GameName)
//			}
//		})
//	}
//}
//
//func TestUserService_AcceptInvite(t *testing.T) {
//	slotId := primitive.NewObjectID()
//
//	tests := []struct {
//		name                   string
//		mockSlotRetrievalError error
//		mockSlot               *entities.Slot
//		mockGameRetrievalError error
//		mockGame               *entities.Game
//		mockBookSlotError      error
//		mockUserRepoError      error
//		expectedError          bool
//	}{
//		{
//			name:                   "Successful Acceptance",
//			mockSlotRetrievalError: nil,
//			mockSlot: &entities.Slot{
//				ID:     slotId,
//				GameID: primitive.NewObjectID(),
//			},
//			mockGameRetrievalError: nil,
//			mockGame: &entities.Game{
//				ID:   primitive.NewObjectID(),
//				Name: "Test Game",
//			},
//			mockBookSlotError: nil,
//			mockUserRepoError: nil,
//			expectedError:     false,
//		},
//		{
//			name:                   "Slot Retrieval Error",
//			mockSlotRetrievalError: errors.New("slot not found"),
//			mockSlot:               nil,
//			mockGameRetrievalError: nil,
//			mockGame:               nil,
//			mockBookSlotError:      nil,
//			mockUserRepoError:      nil,
//			expectedError:          true,
//		},
//		{
//			name:                   "Game Retrieval Error",
//			mockSlotRetrievalError: nil,
//			mockSlot: &entities.Slot{
//				ID:     slotId,
//				GameID: primitive.NewObjectID(),
//			},
//			mockGameRetrievalError: errors.New("game not found"),
//			mockGame:               nil,
//			mockBookSlotError:      nil,
//			mockUserRepoError:      nil,
//			expectedError:          true,
//		},
//		{
//			name:                   "Booking Slot Error",
//			mockSlotRetrievalError: nil,
//			mockSlot: &entities.Slot{
//				ID:     slotId,
//				GameID: primitive.NewObjectID(),
//			},
//			mockGameRetrievalError: nil,
//			mockGame: &entities.Game{
//				ID:   primitive.NewObjectID(),
//				Name: "Test Game",
//			},
//			mockBookSlotError: errors.New("slot booking error"),
//			mockUserRepoError: nil,
//			expectedError:     true,
//		},
//		{
//			name:                   "Delete Invite Error",
//			mockSlotRetrievalError: nil,
//			mockSlot: &entities.Slot{
//				ID:     slotId,
//				GameID: primitive.NewObjectID(),
//			},
//			mockGameRetrievalError: nil,
//			mockGame: &entities.Game{
//				ID:   primitive.NewObjectID(),
//				Name: "Test Game",
//			},
//			mockBookSlotError: nil,
//			mockUserRepoError: errors.New("delete invite error"),
//			expectedError:     true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			// Mock Slot Retrieval
//			mockSlotService.EXPECT().
//				GetSlotById(slotId).
//				Return(tt.mockSlot, tt.mockSlotRetrievalError).
//				Times(1)
//
//			// Only mock Game Retrieval if Slot Retrieval is successful
//			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil {
//				mockGameService.EXPECT().
//					GetGameByID(tt.mockSlot.GameID).
//					Return(tt.mockGame, tt.mockGameRetrievalError).
//					Times(1)
//			}
//
//			// Only mock BookSlot if both Slot and Game Retrievals are successful
//			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil && tt.mockGameRetrievalError == nil && tt.mockGame != nil {
//				mockSlotService.EXPECT().
//					BookSlot(tt.mockGame, tt.mockSlot).
//					Return(tt.mockBookSlotError).
//					Times(1)
//			}
//
//			// Only mock DeleteInvite if Slot booking is successful
//			if tt.mockSlotRetrievalError == nil && tt.mockSlot != nil &&
//				tt.mockGameRetrievalError == nil && tt.mockGame != nil &&
//				tt.mockBookSlotError == nil {
//				mockUserRepo.EXPECT().
//					DeleteInvite(slotId).
//					Return(tt.mockUserRepoError).
//					Times(1)
//			}
//
//			err := userService.AcceptInvite(slotId)
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
//func TestUserService_RejectInvite(t *testing.T) {
//	slotId := primitive.NewObjectID()
//
//	tests := []struct {
//		name          string
//		mockRepoError error
//		expectedError bool
//	}{
//		{
//			name:          "Successful Rejection",
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "Repository Error",
//			mockRepoError: errors.New("error deleting invite"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			mockUserRepo.EXPECT().DeleteInvite(slotId).Return(tt.mockRepoError).Times(1)
//
//			err := userService.RejectInvite(slotId)
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
//func TestUserService_AddResult(t *testing.T) {
//	userId := primitive.NewObjectID()
//
//	tests := []struct {
//		name          string
//		result        string
//		mockRepoError error
//		expectedError bool
//	}{
//		{
//			name:          "Add Win",
//			result:        "win",
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "Add Loss",
//			result:        "loss",
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "Invalid Result",
//			result:        "invalid",
//			mockRepoError: nil,
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			if tt.result == "win" {
//				mockUserRepo.EXPECT().AddWin(userId).Return(tt.mockRepoError).Times(1)
//			} else if tt.result == "loss" {
//				mockUserRepo.EXPECT().AddLoss(userId).Return(tt.mockRepoError).Times(1)
//			}
//
//			err := userService.AddResult(userId, tt.result)
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
//func TestUserService_GetAllUsersByScore(t *testing.T) {
//	tests := []struct {
//		name          string
//		mockUsers     []entities.User
//		mockRepoError error
//		expectedError bool
//		expectedUsers []entities.User
//	}{
//		{
//			name: "Successful Retrieval",
//			mockUsers: []entities.User{
//				{Email: "test1@example.com"},
//				{Email: "test2@example.com"},
//			},
//			mockRepoError: nil,
//			expectedError: false,
//			expectedUsers: []entities.User{
//				{Email: "test1@example.com"},
//				{Email: "test2@example.com"},
//			},
//		},
//		{
//			name:          "Repository Error",
//			mockUsers:     nil,
//			mockRepoError: errors.New("repository error"),
//			expectedError: true,
//			expectedUsers: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			mockUserRepo.EXPECT().GetAllUsersByScore().Return(tt.mockUsers, tt.mockRepoError).Times(1)
//
//			users, err := userService.GetAllUsersByScore()
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.ElementsMatch(t, tt.expectedUsers, users)
//			}
//		})
//	}
//}
