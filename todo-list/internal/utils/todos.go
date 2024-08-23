package utils

import (
	"fmt"
	"project/internal/models"
	"reflect"
)

func GetUserTodos(username string) []string {
	return TodosMap[username].TodoList
}
func AddTodoForUser(username, todo string) error {

	userTodos, exists := TodosMap[username]
	if !exists {
		return fmt.Errorf("user not found")
	}

	userTodos.TodoList = append(userTodos.TodoList, todo)
	TodosMap[username] = userTodos
	return nil
}

func DeleteTodoForUser(username string, index int) error {
	userTodos, exists := TodosMap[username]
	if !exists {
		return fmt.Errorf("user not found")
	}
	userTodos.TodoList = append(userTodos.TodoList[:index], userTodos.TodoList[index+1:]...)
	TodosMap[username] = userTodos
	return nil
}

func MarkTodoAsDoneForUser(username string, index int) error {
	userTodos, exists := TodosMap[username]
	if !exists {
		return fmt.Errorf("user not found")
	}
	doneTodo := userTodos.TodoList[index]
	_ = AddToDailyStatus(username, doneTodo)
	userTodos = TodosMap[username]
	userTodos.TodoList = append(userTodos.TodoList[:index], userTodos.TodoList[index+1:]...)
	TodosMap[username] = userTodos

	return nil
}

func LoadUserTodos(filename string) {
	todoChan := make(chan any)

	go RfileGeneral(todoChan, filename, reflect.TypeOf(models.UserTodos{}))

	for todo := range todoChan {
		todoData, ok := todo.(*models.UserTodos)
		if !ok {
			fmt.Println("Error: received data is not of type models.UserTodos")
		}
		TodosMap[todoData.Username] = todoData.Details
	}
}
