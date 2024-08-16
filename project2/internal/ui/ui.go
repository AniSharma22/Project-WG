package ui

import (
	"project2/internal/app/services"
)

// UI struct holds the UserService and other dependencies
type UI struct {
	userService *services.UserService
}

// NewUI initializes the UI with the provided services
func NewUI(userService *services.UserService) *UI {
	return &UI{
		userService: userService,
	}
}
