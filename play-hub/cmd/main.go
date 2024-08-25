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
	"project2/pkg/utils"
	"syscall"
	"time"
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
	gameService := services.NewGameService(gameRepo)
	slotService := services.NewSlotService(slotRepo, userRepo, gameHistoryRepo)
	userService := services.NewUserService(userRepo, slotService, gameService)
	gameHistoryService := services.NewGameHistoryService(gameHistoryRepo, userService, slotService)
	leaderboardService := services.NewLeaderboardService(leaderboardRepo, userService)
	notificationService := services.NewNotificationService(notificationRepo)
	// Insert today's slots
	err = utils.InsertAllSlots(slotRepo, gameRepo)
	if err != nil {
		log.Fatal(err)
	}

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
}

func InitClient(uri string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	globals.Client = client
	globals.IstLocation, err = time.LoadLocation("Asia/Kolkata")
	return nil
}
