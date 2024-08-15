package entities

type Game struct {
	GameId     string `json:"gameId"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"maxPlayers"`
}
