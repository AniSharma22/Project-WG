package service_test

import (
	service_interfaces "project2/internal/domain/interfaces/service"
	"testing"

	"github.com/golang/mock/gomock"
	"project2/internal/app/services"
	mock_interfaces "project2/tests/mocks/repository"
	mock_services "project2/tests/mocks/service"
)

var (
	ctrl *gomock.Controller

	mockUserRepo         *mock_interfaces.MockUserRepository
	mockSlotRepo         *mock_interfaces.MockSlotRepository
	mockGameRepo         *mock_interfaces.MockGameRepository
	mockLeaderboardRepo  *mock_interfaces.MockLeaderboardRepository
	mockInvitationRepo   *mock_interfaces.MockInvitationRepository
	mockBookingRepo      *mock_interfaces.MockBookingRepository
	mockNotificationRepo *mock_interfaces.MockNotificationRepository

	mockUserService         *mock_services.MockUserService
	mockSlotService         *mock_services.MockSlotService
	mockGameService         *mock_services.MockGameService
	mockLeaderboardService  *mock_services.MockLeaderboardService
	mockInvitationService   *mock_services.MockInvitationService
	mockBookingService      *mock_services.MockBookingService
	mockNotificationService *mock_services.MockNotificationService

	userService         service_interfaces.UserService
	slotService         service_interfaces.SlotService
	gameService         service_interfaces.GameService
	leaderboardService  service_interfaces.LeaderboardService
	invitationService   service_interfaces.InvitationService
	bookingService      service_interfaces.BookingService
	notificationService service_interfaces.NotificationService
)

func setup(t *testing.T) func() {
	// Set up the go mock controller
	ctrl = gomock.NewController(t)

	// Create mock repositories
	mockUserRepo = mock_interfaces.NewMockUserRepository(ctrl)
	mockSlotRepo = mock_interfaces.NewMockSlotRepository(ctrl)
	mockGameRepo = mock_interfaces.NewMockGameRepository(ctrl)
	mockLeaderboardRepo = mock_interfaces.NewMockLeaderboardRepository(ctrl)
	mockInvitationRepo = mock_interfaces.NewMockInvitationRepository(ctrl)
	mockBookingRepo = mock_interfaces.NewMockBookingRepository(ctrl)
	mockNotificationRepo = mock_interfaces.NewMockNotificationRepository(ctrl)

	// Create mock services
	mockUserService = mock_services.NewMockUserService(ctrl)
	mockSlotService = mock_services.NewMockSlotService(ctrl)
	mockGameService = mock_services.NewMockGameService(ctrl)
	mockLeaderboardService = mock_services.NewMockLeaderboardService(ctrl)
	mockInvitationService = mock_services.NewMockInvitationService(ctrl)
	mockBookingService = mock_services.NewMockBookingService(ctrl)
	mockNotificationService = mock_services.NewMockNotificationService(ctrl)

	// Create genuine services
	userService = services.NewUserService(mockUserRepo)
	slotService = services.NewSlotService(mockSlotRepo)
	gameService = services.NewGameService(mockGameRepo)
	bookingService = services.NewBookingService(mockBookingRepo, mockSlotService, mockGameService)
	leaderboardService = services.NewLeaderboardService(mockLeaderboardRepo, mockBookingService)
	invitationService = services.NewInvitationService(mockInvitationRepo, mockBookingService, mockSlotService)
	notificationService = services.NewNotificationService(mockNotificationRepo)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}
