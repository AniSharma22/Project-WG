package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/internal/ui"
	"project2/pkg/globals"
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
	gameHistoryService := services.NewGameHistoryService(gameHistoryRepo, userService)
	leaderboardService := services.NewLeaderboardService(leaderboardRepo, userService)
	notificationService := services.NewNotificationService(notificationRepo)
	// Insert today's slots
	err = insertAllSlots(slotRepo, gameRepo)
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
	return nil
}

func insertAllSlots(slotRepo interfaces.SlotRepository, gameRepo interfaces.GameRepository) error {
	today := time.Now().Format("2006-01-02")

	// Fetch all games
	games, err := gameRepo.GetAllGames()
	if err != nil {
		return fmt.Errorf("error fetching games: %w", err)
	}

	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)

	for _, game := range games {
		// Check for existing slots for this game on today's date
		existingSlots, err := slotRepo.GetSlotsByDate(today, game.ID)
		if err != nil {
			return fmt.Errorf("error checking existing slots for game %s: %w", game.Name, err)
		}

		// If no slots exist, create new slots
		if len(existingSlots) == 0 {
			for current := startTime; current.Before(endTime); current = current.Add(20 * time.Minute) {
				slotEndTime := current.Add(20 * time.Minute)
				if slotEndTime.After(endTime) {
					slotEndTime = endTime
				}

				newSlot := entities.Slot{
					ID:          primitive.NewObjectID(),
					GameID:      game.ID,
					Date:        today,
					StartTime:   current.Format("15:04"),
					EndTime:     slotEndTime.Format("15:04"),
					BookedUsers: []primitive.ObjectID{},
					Results:     []entities.Result{},
				}

				// Insert the new slot
				if _, err := slotRepo.InsertSlot(newSlot); err != nil {
					return fmt.Errorf("error inserting slot for game %s: %w", game.Name, err)
				}
			}
		}
	}

	return nil
}
