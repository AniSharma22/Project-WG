package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"project/internal/config"
	"project/internal/models"
	"project/internal/utils"
	"reflect"
	"strconv"
)

func init() {
	// Check if the users.json file exists and load users if it does
	_, err := os.Stat(config.UsersFile)
	if err == nil {
		go utils.LoadUsers(config.UsersFile)
	} else {
		utils.UserDataLoaded = true
	}

	if _, err := os.Stat(config.ProgressFile); err == nil {
		go utils.LoadUserProgress(config.ProgressFile)
	}

	if _, err := os.Stat(config.TodosFile); err == nil {
		go utils.LoadUserTodos(config.TodosFile)
	}

	// Start loading the course outline
	go utils.LoadCourseOutline(config.CoursesFile)

}

func main() {

	r := mux.NewRouter()

	// Define the routes and methods

	// 1. To return the user todos
	r.HandleFunc("/todos/{username}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(utils.GetUserTodos(username))
		return
	}).Methods("GET")

	// 2. To add a new user todos
	r.HandleFunc("/todos/{username}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Extract the todo from the request body
		var requestBody struct {
			Todo string `json:"todo"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Call the AddTodoForUser function
		err = utils.AddTodoForUser(username, requestBody.Todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WfileGeneral(config.TodosFile, utils.TodosMap, reflect.TypeOf(models.UserTodos{}))
		// Respond with success
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Todo added successfully!")
	}).Methods("POST")

	r.HandleFunc("/todos/{username}/{index}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		indexStr := vars["index"]

		// Convert the index from string to int
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		// Call the DeleteTodoForUser function
		err = utils.DeleteTodoForUser(username, index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WfileGeneral(config.TodosFile, utils.TodosMap, reflect.TypeOf(models.UserTodos{}))
		// Respond with success message
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Todo deleted successfully!")
	}).Methods("DELETE")

	r.HandleFunc("/todos/{username}/{index}/done", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		indexStr := vars["index"]

		// Convert the index from string to int
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		// Call the MarkTodoAsDoneForUser function
		err = utils.MarkTodoAsDoneForUser(username, index)
		if err != nil {
			if err.Error() == "user not found" {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		utils.WfileGeneral(config.TodosFile, utils.TodosMap, reflect.TypeOf(models.UserTodos{}))
		// Respond with success message
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Todo marked as done successfully!")
	}).Methods("PUT")

	fmt.Println("server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8081", r))

	//for {
	//	var choice string
	//	fmt.Println("\033[1;36m") // Cyan bold
	//	fmt.Println("===================================")
	//	fmt.Println("     		  WELCOME    ")
	//	fmt.Println("===================================")
	//	fmt.Println("\033[0m") // Reset color
	//	fmt.Println("Please choose an option:")
	//	fmt.Println("1. Signup")
	//	fmt.Println("2. Login")
	//	fmt.Println("3. Exit")
	//
	//	fmt.Print("Enter your choice (1, 2 or 3): ")
	//	_, err := fmt.Scanln(&choice)
	//	if err != nil {
	//		fmt.Println("Error reading input:", err)
	//		continue
	//	}
	//
	//	switch choice {
	//	case "1":
	//		fmt.Println("------------")
	//		fmt.Println("Signup")
	//		fmt.Println("------------")
	//		ui.HandleUserAction("signup")
	//	case "2":
	//		fmt.Println("------------")
	//		fmt.Println("Login")
	//		fmt.Println("------------")
	//		ui.HandleUserAction("login")
	//	case "3":
	//		if utils.ProgressMap != nil {
	//			utils.WfileGeneral(config.ProgressFile, utils.ProgressMap, reflect.TypeOf(models.UserProgress{}))
	//		}
	//		if utils.TodosMap != nil {
	//			utils.WfileGeneral(config.TodosFile, utils.TodosMap, reflect.TypeOf(models.UserTodos{}))
	//		}
	//		fmt.Println("Exiting...")
	//
	//		return
	//	default:
	//		fmt.Println("Invalid choice. Please enter 1 for Signup, 2 for Login, or 3 for Exit.")
	//	}
	//}
}
