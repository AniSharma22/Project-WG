package entities

type Result struct {
	ResultId    string `json:"result_id"`
	SlotId      string `json:"slot_id"`
	GameId      string `json:"game_id"`
	WinningUser []User `json:"winning_user"`
	LosingUser  []User `json:"losing_user"`
	Score       string `json:"score"`
}
