package models

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CourseOutline struct {
	Modules []string
}
