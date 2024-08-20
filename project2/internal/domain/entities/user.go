package entities

type User struct {
	UserId     string      `json:"user_id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	PhoneNo    string      `json:"phoneNo"`
	Gender     string      `json:"gender"`
	GameStats  []GameStats `json:"game_stats"`
	TotalWins  int         `json:"totalWins"`
	TotalLoss  int         `json:"totalLoss"`
	TotalGames int         `json:"totalGames"`
	Score      float32     `json:"totalScore"`
	Role       string      `json:"role"`
}

type GameStats struct {
	GameID     string  `json:"game_id"`
	Wins       int     `json:"wins"`
	Losses     int     `json:"losses"`
	TotalGames int     `json:"total_games"`
	Score      float32 `json:"score"`
}
