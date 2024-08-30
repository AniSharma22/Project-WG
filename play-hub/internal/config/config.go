package config

// DBConfig holds the database-related configuration
type DBConfig struct {
	DBURI                   string
	DBName                  string
	UsersCollection         string
	NotificationsCollection string
	LeaderboardsCollection  string
	SlotsCollection         string
	GameHistoryCollection   string
	GamesCollection         string

	// Add other collections here
}

var DB = DBConfig{
	DBURI:                   "mongodb://localhost:27017",
	DBName:                  "play-hub",
	UsersCollection:         "Users",
	NotificationsCollection: "Notifications",
	LeaderboardsCollection:  "Leaderboards",
	SlotsCollection:         "Slots",
	GameHistoryCollection:   "GameHistory",
	GamesCollection:         "Games",
}
