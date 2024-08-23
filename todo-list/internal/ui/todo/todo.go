package todo

import (
	"bufio"
	"fmt"
	"os"
	"project/internal/utils"
	"strconv"
	"strings"
)

// ManageTodos handles t.odo-related actions
func ManageTodos(username string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Todo Management")
		fmt.Println("1. View Todos")
		fmt.Println("2. Add Todo")
		fmt.Println("3. Delete Todo")
		fmt.Println("4. Mark Todo as Done")
		fmt.Println("5. Go Back")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			viewTodos(username)

		case "2":
			addTodo(username)

		case "3":
			deleteTodo(username)

		case "4":
			markTodoAsDone(username)

		case "5":
			return // Go back (exit the function)

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// viewTodos displays the list of todos for the user
func viewTodos(username string) {
	todos := utils.GetUserTodos(username)
	fmt.Print("\n")
	fmt.Println("------------------")
	fmt.Println("Your Todos:")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}
	fmt.Println("------------------")
}

// addTodo allows the user to add a new t.odo
func addTodo(username string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the new todo: ")
	todo, _ := reader.ReadString('\n')
	todo = strings.TrimSpace(todo)
	if err := utils.AddTodoForUser(username, todo); err != nil {
		fmt.Println("Error adding todo:", err)
	} else {
		fmt.Println("Todo added successfully!")
	}
}

// deleteTodo allows the user to delete a t.odo by index
func deleteTodo(username string) {
	if len(utils.GetUserTodos(username)) == 0 {
		fmt.Println("You have no todos")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	viewTodos(username)
	fmt.Print("Enter the number of the todo to delete: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(utils.GetUserTodos(username)) {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if err := utils.DeleteTodoForUser(username, index-1); err != nil {
		fmt.Println("Error deleting todo:", err)
	} else {
		fmt.Println("Todo deleted successfully!")
	}
}

// markTodoAsDone allows the user to mark a t.odo as done by index
func markTodoAsDone(username string) {
	if len(utils.GetUserTodos(username)) == 0 {
		fmt.Println("You have no todos")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	viewTodos(username)
	fmt.Print("Enter the number of the todo to mark as done: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(utils.GetUserTodos(username)) {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if err := utils.MarkTodoAsDoneForUser(username, index-1); err != nil {
		fmt.Println("Error marking todo as done:", err)
	} else {
		fmt.Println("Todo marked as done!")
	}
}
