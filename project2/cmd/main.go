package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/ui"
	"project2/pkg/globals"
	"syscall"
)

func main() {
	err := InitClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize all repositories
	userRepo := repositories.NewUserRepo()
	gameRepo := repositories.NewGameRepo()
	slotRepo := repositories.NewSlotRepo()
	gameHistoryRepo := repositories.NewGameHistoryRepo()
	leaderboardRepo := repositories.NewLeaderboardRepo()
	notificationRepo := repositories.NewNotificationRepo()

	// Initialize all services
	userService := services.NewUserService(userRepo)
	gameService := services.NewGameService(gameRepo)
	slotService := services.NewSlotService(slotRepo)
	gameHistoryService := services.NewGameHistoryService(gameHistoryRepo)
	leaderboardService := services.NewLeaderboardService(leaderboardRepo)
	notificationService := services.NewNotificationService(notificationRepo)

	// Handling graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Graceful shutdown initiated...")
		// Wait for all operations to complete if necessary
		fmt.Println("All operations completed. Exiting.")
		os.Exit(0)
	}()

	// Set up the UI with all services
	appUI := ui.NewUI(userService, gameService, slotService, gameHistoryService, leaderboardService, notificationService)

	appUI.ShowMainMenu()

	// Waiting for all operations to complete before exiting
	// Consider if this is necessary or remove if not needed
	// userService.WaitForCompletion()
}

func InitClient(uri string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	globals.Client = client
	return nil
}
