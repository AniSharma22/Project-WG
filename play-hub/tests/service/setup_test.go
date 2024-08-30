package service_test

import (
	"github.com/golang/mock/gomock"
	"project2/internal/app/services"
	"project2/internal/domain/interfaces"
	mock_interfaces "project2/tests/mocks/repository"
	mock_services "project2/tests/mocks/service"
	"testing"
)

var (
	ctrl                    *gomock.Controller
	mockUserRepo            *mock_interfaces.MockUserRepository
	mockSlotRepo            *mock_interfaces.MockSlotRepository
	mockGameRepo            *mock_interfaces.MockGameRepository
	mockGameHistoryRepo     *mock_interfaces.MockGameHistoryRepository
	mockLeaderboardRepo     *mock_interfaces.MockLeaderboardRepository
	mockUserService         *mock_services.MockUserService
	mockSlotService         *mock_services.MockSlotService
	mockGameService         *mock_services.MockGameService
	mockGameHistoryService  *mock_services.MockGameHistoryService
	mockLeaderboardService  *mock_services.MockLeaderboardService
	mockNotificationService *mock_services.MockNotificationService
	userService             interfaces.UserService
	slotService             interfaces.SlotService
	gameService             interfaces.GameService
	gameHistoryService      interfaces.GameHistoryService
	leaderboardService      interfaces.LeaderboardService
	notificationService     interfaces.NotificationService
)

func setup(t *testing.T) func() {
	// Set up the go mock controller
	ctrl = gomock.NewController(t)

	// Create mock repositories
	mockUserRepo = mock_interfaces.NewMockUserRepository(ctrl)
	mockGameRepo = mock_interfaces.NewMockGameRepository(ctrl)
	mockLeaderboardRepo = mock_interfaces.NewMockLeaderboardRepository(ctrl)
	mockGameHistoryRepo = mock_interfaces.NewMockGameHistoryRepository(ctrl)
	mockSlotRepo = mock_interfaces.NewMockSlotRepository(ctrl)

	// Create mock services
	mockUserService = mock_services.NewMockUserService(ctrl)
	mockSlotService = mock_services.NewMockSlotService(ctrl)
	mockGameService = mock_services.NewMockGameService(ctrl)
	mockGameHistoryService = mock_services.NewMockGameHistoryService(ctrl)
	mockLeaderboardService = mock_services.NewMockLeaderboardService(ctrl)
	mockNotificationService = mock_services.NewMockNotificationService(ctrl)

	// Create Genuine Services
	userService = services.NewUserService(mockUserRepo, mockSlotService, mockGameService)
	gameService = services.NewGameService(mockGameRepo)
	leaderboardService = services.NewLeaderboardService(mockLeaderboardRepo, mockUserService)
	gameHistoryService = services.NewGameHistoryService(mockGameHistoryRepo, mockUserService, mockSlotService)
	slotService = services.NewSlotService(mockSlotRepo, mockUserRepo, mockGameHistoryRepo)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}
