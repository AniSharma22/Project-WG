package ui

import (
	"bufio"
	"project2/internal/domain/interfaces"
)

// UI struct holds the UserService, bufio.Reader, and other dependencies
type UI struct {
	userService         interfaces.UserService
	gameService         interfaces.GameService
	slotService         interfaces.SlotService
	gameHistoryService  interfaces.GameHistoryService
	leaderboardService  interfaces.LeaderboardService
	notificationService interfaces.NotificationService
	reader              *bufio.Reader
}

// NewUI initializes the UI with the provided services and a bufio.Reader
func NewUI(userService interfaces.UserService, gameService interfaces.GameService, slotService interfaces.SlotService, gameHistoryService interfaces.GameHistoryService, leaderboardService interfaces.LeaderboardService, notificationService interfaces.NotificationService, reader *bufio.Reader) *UI {
	return &UI{
		userService:         userService,
		gameService:         gameService,
		slotService:         slotService,
		gameHistoryService:  gameHistoryService,
		leaderboardService:  leaderboardService,
		notificationService: notificationService,
		reader:              reader,
	}
}
