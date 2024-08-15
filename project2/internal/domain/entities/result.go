package entities

type Result struct {
	ResultId    string `json:"resultId"`
	SlotId      string `json:"slotId"`
	WinningUser User   `json:"winningUser"`
	LosingUser  User   `json:"losingUser"`
	Score       string `json:"score"`
}
