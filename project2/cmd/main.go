package main

import (
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/ui"
)

func main() {

	// Initialize the repository
	userRepo := repositories.NewUserRepo() // Ensure you have a constructor function

	// Initialize the service
	userService := services.NewUserService(userRepo)

	// set up the UI with user service
	appUI := ui.NewUI(userService)

	appUI.ShowMainMenu()
}
