package main

import (
	"fmt"
	"os"
	"os/signal"
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/ui"
	"syscall"
)

func main() {

	// Initialize the repository
	userRepo := repositories.NewUserRepo()
	gameRepo := repositories.NewGameRepo()

	// Initialize the service
	userService := services.NewUserService(userRepo)
	gameService := services.NewGameService(gameRepo)

	// handling graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("graceful shutdown initiated...")
		userService.WaitForCompletion()
		fmt.Println("All operations completed. Exiting.")
		os.Exit(0)
	}()

	// set up the UI with user service
	appUI := ui.NewUI(userService, gameService)

	appUI.ShowMainMenu()

	// waiting for all operations to complete before exiting
	userService.WaitForCompletion()
}
