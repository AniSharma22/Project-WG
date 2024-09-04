package ui

import (
	"bufio"
	"project2/internal/domain/interfaces/service"
)

// UI struct holds the UserService, bufio.Reader, and other dependencies
type UI struct {
	userService         service_interfaces.UserService
	gameService         service_interfaces.GameService
	slotService         service_interfaces.SlotService
	bookingService      service_interfaces.BookingService
	invitationService   service_interfaces.InvitationService
	leaderboardService  service_interfaces.LeaderboardService
	notificationService service_interfaces.NotificationService
	reader              *bufio.Reader
}

// NewUI initializes the UI with the provided services and a bufio.Reader
func NewUI(userService service_interfaces.UserService, gameService service_interfaces.GameService, slotService service_interfaces.SlotService, bookingService service_interfaces.BookingService, invitationService service_interfaces.InvitationService, leaderboardService service_interfaces.LeaderboardService, notificationService service_interfaces.NotificationService, reader *bufio.Reader) *UI {
	return &UI{
		userService:         userService,
		gameService:         gameService,
		slotService:         slotService,
		bookingService:      bookingService,
		invitationService:   invitationService,
		leaderboardService:  leaderboardService,
		notificationService: notificationService,
		reader:              reader,
	}
}
