package models

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Module represents a single module in a course
type Module struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Course represents a course with multiple modules
type Course struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Modules []Module `json:"modules"`
}

// UserProgress stores the progress of multiple users
type UserProgress struct {
	Username         string `json:"username"`
	CompletedModules []int  `json:"completedModules"`
}
