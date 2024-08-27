package ui

import (
	"bufio"
	"bytes"
	"fmt"
	"project2/internal/domain/interfaces"
	"project2/internal/ui"
	"testing"

	"github.com/golang/mock/gomock"
	mock_services "project2/tests/mocks/service"
)

var (
	ctrl                    *gomock.Controller
	mockUserService         interfaces.UserService
	mockGameService         interfaces.GameService
	mockSlotService         interfaces.SlotService
	mockGameHistoryService  interfaces.GameHistoryService
	mockLeaderboardService  interfaces.LeaderboardService
	mockNotificationService interfaces.NotificationService
	testUI                  *ui.UI
)

func setup(t *testing.T) func() {
	ctrl = gomock.NewController(t)
	mockUserService = mock_services.NewMockUserService()
	mockGameService = mock_services.NewMockGameService()
	mockSlotService = mock_services.NewMockSlotService()
	mockGameHistoryService = mock_services.NewMockGameHistoryService()
	mockLeaderboardService = mock_services.NewMockLeaderboardService()
	mockNotificationService = mock_services.NewMockNotificationService()

	testUI := ui.NewUI(mockUserService, mockGameService, mockSlotService, mockGameHistoryService, mockLeaderboardService, mockNotificationService, nil)
	fmt.Println(testUI)
	return func() {
		ctrl.Finish()
	}
}

func TestShowMainMenu_Signup(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Simulate user input for selecting "Signup"
	input := "1\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))

	// Reinitialize testUI with the mocked reader
	testUI = ui.NewUI(mockUserService, mockGameService, mockSlotService, mockGameHistoryService, mockLeaderboardService, mockNotificationService, reader)

	// Mock the expected behavior for ShowSignupPage
	mockUserService.EXPECT().Signup(gomock.Any()).Times(1)

	// Call the function to test
	testUI.ShowMainMenu()

	// No specific assertions needed, the expectation on Signup being called is enough.
}

func TestShowMainMenu_Login(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Simulate user input for selecting "Login"
	input := "2\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))

	// Reinitialize testUI with the mocked reader
	testUI = ui.NewUI(mockUserService, mockGameService, mockSlotService, mockGameHistoryService, mockLeaderboardService, mockNotificationService, reader)

	// Mock the expected behavior for ShowLoginPage
	mockUserService.EXPECT().Login(gomock.Any(), gomock.Any()).Times(1)

	// Call the function to test
	testUI.ShowMainMenu()

	// No specific assertions needed, the expectation on Login being called is enough.
}

func TestShowMainMenu_Exit(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Simulate user input for selecting "Exit"
	input := "3\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))

	// Reinitialize testUI with the mocked reader
	testUI = ui.NewUI(mockUserService, mockGameService, mockSlotService, mockGameHistoryService, mockLeaderboardService, mockNotificationService, reader)

	// Call the function to test
	testUI.ShowMainMenu()

	// No specific assertions needed, just ensuring the function returns.
}
