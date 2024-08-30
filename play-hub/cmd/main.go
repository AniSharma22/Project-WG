package main

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/config"
	"project2/internal/ui"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"sync"
	"syscall"
	"time"
)

// Declare a variable to hold the singleton instance of the MongoDB client
var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// InitClient initializes a MongoDB client with a connection pool of 30
func InitClient(uri string) (*mongo.Client, error) {
	var err error

	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		clientOptions.SetMaxPoolSize(30) // Set the maximum number of connections in the pool

		clientInstance, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		globals.IstLocation, err = time.LoadLocation("Asia/Kolkata")
		if err != nil {
			log.Fatalf("Failed to load location: %v", err)
		}
	})

	return clientInstance, err
}

func main() {
	client, err := InitClient(config.DB.DBURI)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize all repositories with the MongoDB client
	userRepo := repositories.NewUserRepo(client)
	gameRepo := repositories.NewGameRepo(client)
	slotRepo := repositories.NewSlotRepo(client)
	gameHistoryRepo := repositories.NewGameHistoryRepo(client)
	leaderboardRepo := repositories.NewLeaderboardRepo(client)
	notificationRepo := repositories.NewNotificationRepo(client)

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
	appUI := ui.NewUI(userService, gameService, slotService, gameHistoryService, leaderboardService, notificationService, bufio.NewReader(os.Stdin))
	appUI.ShowMainMenu()
}
