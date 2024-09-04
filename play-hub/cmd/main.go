package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"project2/internal/app/repositories"
	"project2/internal/app/services"
	"project2/internal/config"
	"project2/internal/ui"
	"project2/pkg/utils"
	"syscall"
)

// InitClient initializes a PostgresSQL client
func InitClient() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}

	return db, err
}

func main() {

	// Initialize PostgresSQL client
	client, err := InitClient()
	if err != nil {
		log.Fatal("Error initializing PostgresSQL client:", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal("Error closing PostgresSQL client:", err)
		}
	}()

	// Initialize repositories
	userRepo := repositories.NewUserRepo(client)
	gameRepo := repositories.NewGameRepo(client)
	slotRepo := repositories.NewSlotRepo(client)
	invitationRepo := repositories.NewInvitationRepo(client)
	bookingRepo := repositories.NewBookingRepo(client)
	leaderboardRepo := repositories.NewLeaderboardRepo(client)
	notificationRepo := repositories.NewNotificationRepo(client)

	// Initialize services
	gameService := services.NewGameService(gameRepo)
	slotService := services.NewSlotService(slotRepo)
	userService := services.NewUserService(userRepo)
	bookingService := services.NewBookingService(bookingRepo, slotService, gameService)
	invitationService := services.NewInvitationService(invitationRepo, bookingService, slotService)
	leaderboardService := services.NewLeaderboardService(leaderboardRepo, bookingService)
	notificationService := services.NewNotificationService(notificationRepo)

	// Insert today's slots
	err = utils.InsertAllSlots(context.Background(), slotRepo, gameRepo)
	if err != nil {
		log.Fatal("Error inserting slots:", err)
	}

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Graceful shutdown initiated...")
		// Additional cleanup if necessary
		fmt.Println("All operations completed. Exiting.")
		os.Exit(0)
	}()

	// Initialize and display the UI
	appUI := ui.NewUI(userService, gameService, slotService, bookingService, invitationService, leaderboardService, notificationService, bufio.NewReader(os.Stdin))
	appUI.ShowMainMenu()
}
