package ui

import (
	"bufio"
	"os"
	"project2/internal/app/services"
)

// UI struct holds the UserService, bufio.Reader, and other dependencies
type UI struct {
	userService         *services.UserService
	gameService         *services.GameService
	slotService         *services.SlotService
	gameHistoryService  *services.GameHistoryService
	leaderboardService  *services.LeaderboardService
	notificationService *services.NotificationService
	reader              *bufio.Reader
}

// NewUI initializes the UI with the provided services and a bufio.Reader
func NewUI(userService *services.UserService, gameService *services.GameService, slotService *services.SlotService, gameHistoryService *services.GameHistoryService, leaderboardService *services.LeaderboardService, notificationService *services.NotificationService) *UI {
	return &UI{
		userService:         userService,
		gameService:         gameService,
		slotService:         slotService,
		gameHistoryService:  gameHistoryService,
		leaderboardService:  leaderboardService,
		notificationService: notificationService,
		reader:              bufio.NewReader(os.Stdin), // Initialize the reader to read from standard input
	}
}
