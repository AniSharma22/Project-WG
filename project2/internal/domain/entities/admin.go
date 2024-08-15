package entities

type Admin struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phoneNo"`
	Role     string `json:"role"`
}

func (u *Admin) Signup() {

}

func (u *Admin) Login() {

}

func (u *Admin) Logout()     {}
func (u *Admin) AddNewGame() {}

func (u *Admin) DeleteGame() {}

func (u *Admin) GetUserStats() {}
