package entities

type Leaderboard struct {
	GameId   string `json:"gameId"`
	Rankings []User `json:"rankings"`
}
