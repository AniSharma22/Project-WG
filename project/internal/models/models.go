package models

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CourseOutline Course represents a course with multiple modules
type CourseOutline struct {
	ID      int           `json:"id"`
	Details CourseDetails `json:"details"`
}
type CourseDetails struct {
	Title   string   `json:"title"`
	Modules []Module `json:"modules"`
}

// Module represents a single module in a course
type Module struct {
	ID    float64 `json:"id"`
	Title string  `json:"title"`
}

type UserProgress struct {
	Username string             `json:"username"`
	Progress []CompletedSection `json:"progress"`
}

type CompletedSection struct {
	CourseID          int       `json:"courseID"`
	CompletedSections []float64 `json:"completedSections"`
}

type UserTodos struct {
	Username string      `json:"username"`
	Details  UserDetails `json:"details"`
}

type UserDetails struct {
	TodoList    []string     `json:"todo_list"`
	DailyStatus []TaskStatus `json:"daily_status"`
}

// TaskStatus represents the status of a task at a specific time and date.
type TaskStatus struct {
	Time string `json:"time"`
	Date string `json:"date"`
	Task string `json:"task"`
}
