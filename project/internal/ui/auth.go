package ui

import (
	"fmt"
	"project/internal/config"
	"project/internal/models"
	"project/internal/utils"
	"reflect"
)

func createEmptyEntryForUser(username string) {
	// Initialize a CompletedSection instance with course ID 1 and an empty slice of CompletedSections
	initialProgress := models.CompletedSection{
		CourseID:          1,
		CompletedSections: []float64{},
	}

	// Assign the initial progress to the user's entry in ProgressMap
	utils.ProgressMap[username] = []models.CompletedSection{initialProgress}
}

func createEmptyEntryForTodos(username string) {
	initialTodo := models.UserDetails{
		TodoList:    []string{},
		DailyStatus: []models.TaskStatus{},
	}
	utils.TodosMap[username] = initialTodo
}

func Signup(username, password string) {

	hashedPass, _ := utils.HashPass(password)
	utils.UserMap[username] = hashedPass
	go utils.WfileGeneral(config.UsersFile, utils.UserMap, reflect.TypeOf(models.UserData{}))
	createEmptyEntryForUser(username)
	createEmptyEntryForTodos(username)
	if username == "anish" && password == "Anish2003@" {
		HandleAdminDashboard(username)
	}
	HandleUserDashboard(username)
}

func Login(username, password string) {

	if username == "anish" && password == "Anish2003@" {
		fmt.Println("Login successful!")
		HandleAdminDashboard(username)
		return
	}
	hashedPass, _ := utils.HashPass(password)
	if utils.IsUsernameTaken(username) && utils.VerifyPassword(password, hashedPass) {
		fmt.Println("Login successful!")
		HandleUserDashboard(username)
		return
	}
	fmt.Println("Username or Password is wrong!!")
	return
}
