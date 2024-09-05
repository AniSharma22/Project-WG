package scripts

import (
	"database/sql"
	"fmt"
	"log"
	"project2/internal/config"

	_ "github.com/lib/pq"
)

func InitializeTables() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	createTables := []string{
		`CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    mobile_number VARCHAR(15) UNIQUE,
    gender VARCHAR(10) CHECK (gender IN ('male', 'female', 'other')),
    role VARCHAR(10) CHECK (role IN ('user', 'admin')) DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
`,

		`CREATE TABLE IF NOT EXISTS games (
			game_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			game_name VARCHAR(255) NOT NULL,
			min_players INT,
			max_players INT,
			instances INT,
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS slots (
			slot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			game_id UUID REFERENCES games(game_id),
			slot_date DATE NOT NULL,
			start_time TIMESTAMPTZ NOT NULL,
			end_time TIMESTAMPTZ NOT NULL,
			is_booked BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS bookings (
			booking_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			slot_id UUID REFERENCES slots(slot_id),
			user_id UUID REFERENCES users(user_id),
    		result VARCHAR(5) CHECK (result IN ('win', 'loss', 'pending')) DEFAULT 'pending',
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS invitations (
			invitation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			inviting_user_id UUID REFERENCES users(user_id),
			invited_user_id UUID REFERENCES users(user_id),
    		slot_id UUID REFERENCES slots(slot_id),
			status VARCHAR(10) CHECK (status IN ('pending', 'accepted', 'declined')) DEFAULT 'pending',
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS notifications (
			notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(user_id),
			message TEXT NOT NULL,
			is_read BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS leaderboard (
			score_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(user_id),
			game_id UUID REFERENCES games(game_id),
			wins INT DEFAULT 0,
			losses INT DEFAULT 0,
			score FLOAT DEFAULT 0.0,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, table := range createTables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}

	fmt.Println("Tables created successfully!")
}
