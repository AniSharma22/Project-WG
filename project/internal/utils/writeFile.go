package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"project/internal/models"
)

// Wfile writes the contents of userMap to a file.
func Wfile(userMap map[string]string) {
	// Open or create the file for writing
	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Convert map to slice of User structs
	var users []models.UserData
	for username, password := range userMap {
		users = append(users, models.UserData{
			Username: username,
			Password: password,
		})
	}

	// Create a JSON encoder and write the slice to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing

	if err := encoder.Encode(users); err != nil {
		fmt.Println("Error encoding data to file:", err)
	}
}
