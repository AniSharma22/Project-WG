package service_test

import (
	"errors"
	"project2/internal/app/services"
	mock_interfaces "project2/tests/mocks/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain/entities"
)

var (
	ctrl         *gomock.Controller
	mockUserRepo *mock_interfaces.MockUserRepository
	userService  *services.UserService
)

func setup(t *testing.T) func() {
	// Set up the gomock controller
	ctrl = gomock.NewController(t)

	// Create a mock UserRepository
	mockUserRepo = mock_interfaces.NewMockUserRepository(ctrl)

	// Initialize the UserService with the mock repository
	userService = services.NewUserService(mockUserRepo, nil, nil)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}

func TestUserService_Signup(t *testing.T) {
	// Create a table of test cases
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
			// Run setup and defer cleanup
			teardown := setup(t)
			defer teardown()

			// Set up the expected behavior for the mock repository
			mockUserRepo.EXPECT().CreateUser(gomock.AssignableToTypeOf(tt.newUser)).Return(tt.mockRepoError).Times(1)

			// Call the Signup method
			err := userService.Signup(tt.newUser)

			// Assert the results
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
