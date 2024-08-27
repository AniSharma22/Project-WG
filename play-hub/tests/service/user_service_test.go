package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/app/services"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	mock_interfaces "project2/tests/mocks/repository"
	mock_service "project2/tests/mocks/service"
	"testing"
)

var (
	ctrl         *gomock.Controller
	mockUserRepo *mock_interfaces.MockUserRepository
	userService  interfaces.UserService
)

func setup(t *testing.T) func() {
	// Set up the gomock controller
	ctrl = gomock.NewController(t)

	// Create a mock UserRepository
	mockUserRepo = mock_interfaces.NewMockUserRepository(ctrl)

	// Initialize the UserService with the mock repository
	slotService := mock_service.NewMockSlotService()
	gameService := mock_service.NewMockGameService()
	userService = services.NewUserService(mockUserRepo, slotService, gameService)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}

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
				Password: "$2a$14$7NcnNTGL5l0GyjnYloyWQeWc9wT82T3SoKgpNE38xiYGotmnsPH", // Assumes the password check is internal
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

//func TestUserService_GetPendingInvites(t *testing.T) {
//	now := time.Now()
//	slotId := primitive.NewObjectID()
//	invite := models.Invite{
//		SlotId: slotId,
//	}
//
//	tests := []struct {
//		name            string
//		mockInvites     []models.Invite
//		mockSlot        *entities.Slot
//		mockGame        *entities.Game
//		mockRepoError   error
//		expectedError   bool
//		expectedInvites []models.Invite
//	}{
//		{
//			name: "Successful Retrieval",
//			mockInvites: []models.Invite{
//				invite,
//			},
//			mockSlot: &entities.Slot{
//				ID:        slotId,
//				StartTime: now.Add(1 * time.Hour),
//			},
//			mockGame: &entities.Game{
//				ID:   primitive.NewObjectID(),
//				Name: "Test Game",
//			},
//			mockRepoError: nil,
//			expectedError: false,
//			expectedInvites: []models.Invite{
//				{
//					SlotId:      slotId,
//					GameName:    "Test Game",
//					Date:        time.Now().Truncate(24 * time.Hour),
//					StartTime:   now.Add(1 * time.Hour).Format("15:04"),
//					EndTime:     now.Add(1 * time.Hour).Add(20 * time.Minute).Format("15:04"),
//					BookedUsers: []string{},
//				},
//			},
//		},
//		{
//			name:            "Slot Retrieval Error",
//			mockInvites:     []models.Invite{invite},
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
//			for _, invite := range tt.mockInvites {
//				mockUserRepo.EXPECT().GetPendingInvites(globals.ActiveUser).Return(tt.mockInvites, tt.mockRepoError).Times(1)
//				if tt.mockSlot != nil {
//					mockSlotService.EXPECT().GetSlotById(invite.SlotId).Return(tt.mockSlot, nil).Times(1)
//				}
//				if tt.mockGame != nil {
//					mockGameService.EXPECT().GetGameByID(tt.mockSlot.GameID).Return(tt.mockGame, nil).Times(1)
//				}
//			}
//
//			invites, err := userService.GetPendingInvites()
//
//			if tt.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.ElementsMatch(t, tt.expectedInvites, invites)
//			}
//		})
//	}
//}

//func TestUserService_AcceptInvite(t *testing.T) {
//	slotId := primitive.NewObjectID()
//
//	tests := []struct {
//		name          string
//		mockSlot      *entities.Slot
//		mockGame      *entities.Game
//		mockRepoError error
//		expectedError bool
//	}{
//		{
//			name: "Successful Acceptance",
//			mockSlot: &entities.Slot{
//				ID:          slotId,
//				GameID:      primitive.NewObjectID(),
//				BookedUsers: []primitive.ObjectID{},
//			},
//			mockGame: &entities.Game{
//				ID:   primitive.NewObjectID(),
//				Name: "Test Game",
//			},
//			mockRepoError: nil,
//			expectedError: false,
//		},
//		{
//			name:          "Slot Retrieval Error",
//			mockSlot:      nil,
//			mockGame:      nil,
//			mockRepoError: errors.New("slot not found"),
//			expectedError: true,
//		},
//		{
//			name:          "Game Retrieval Error",
//			mockSlot:      &entities.Slot{ID: slotId},
//			mockGame:      nil,
//			mockRepoError: errors.New("game not found"),
//			expectedError: true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			teardown := setup(t)
//			defer teardown()
//
//			mockSlotService.EXPECT().GetSlotById(slotId).Return(tt.mockSlot, nil).Times(1)
//			if tt.mockSlot != nil {
//				mockGameService.EXPECT().GetGameByID(tt.mockSlot.GameID).Return(tt.mockGame, nil).Times(1)
//			}
//			if tt.mockSlot != nil && tt.mockGame != nil {
//				mockSlotService.EXPECT().BookSlot(tt.mockGame, tt.mockSlot).Return(nil).Times(1)
//			}
//			mockUserRepo.EXPECT().DeleteInvite(slotId).Return(tt.mockRepoError).Times(1)
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
