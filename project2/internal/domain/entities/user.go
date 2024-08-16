package entities

type User struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	PhoneNo    string `json:"phoneNo"`
	Gender     string `json:"gender"`
	TotalWins  int    `json:"totalWins"`
	TotalLoss  int    `json:"totalLoss"`
	TotalGames int    `json:"totalGames"`
	Role       string `json:"role"`
}

func (u *User) Signup() {

}

func (u *User) Login() {

}
func (u *User) Logout() {}
func (u *User) ViewStats() {
}

func (u *User) BookSlot(slot Slot) {

}