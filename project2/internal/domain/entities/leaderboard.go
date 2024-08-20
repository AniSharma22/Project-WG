package entities

type Leaderboard struct {
	OverallRankings []User        `json:"overallRankings"` // Overall leaderboard
	GameRankings    []GameRanking `json:"gameRankings"`    // Game-specific leaderboards
}

type GameRanking struct {
	GameId   string `json:"gameId"`
	Rankings []User `json:"rankings"`
}
