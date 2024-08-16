package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"project2/internal/config"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"project2/pkg/utils"
	"reflect"
)

func init() {
	if _, err := os.Stat(config.UsersFile); err == nil {
		go loadAllUsers()
	}
}

type userRepo struct {
	// Database connection or other dependencies
}

func NewUserRepo() interfaces.UserRepository {
	return &userRepo{}
}

func (r *userRepo) CreateUser(user *entities.User) error {
	// Read the existing users from the file
	file, err := os.OpenFile(config.UsersFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var users []entities.User

	// Decode existing users from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing users:", err)
		return err
	}

	// Append the new user to the users array
	users = append(users, *user)

	// Truncate the file to overwrite it with the updated users array
	if err := file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}

	// Move the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return err
	}

	// Encode the updated users array to JSON and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(users); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	return nil
}

func loadAllUsers() {
	userDataChan := make(chan any)
	go utils.StreamJSONObjects(userDataChan, config.UsersFile, reflect.TypeOf(entities.User{}))

	for user := range userDataChan {
		// Type assertion
		userData, ok := user.(*entities.User)
		if !ok {
			fmt.Println("Error: received data is not of type models.UserData")
			continue
		}

		globals.UsersMap[userData.Email] = *userData
	}
}
