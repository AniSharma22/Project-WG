package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"project/internal/models"
)

// Rfile reads a JSON array from a file and sends user data to the channel.
func Rfile(userDataChan chan<- models.UserData) {
	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		close(userDataChan)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	// Read the opening bracket of the JSON array
	_, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading JSON array start:", err)
		close(userDataChan)
		return
	}

	for decoder.More() {
		var user models.UserData
		// Decode the next JSON object
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Println("Error decoding JSON object:", err)
			continue
		}
		userDataChan <- user
	}

	// Read the closing bracket of the JSON array
	_, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading JSON array end:", err)
	}

	close(userDataChan) // Close the channel when done
}
