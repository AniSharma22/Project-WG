package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"project/internal/models"
	"reflect"
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

func RfileGeneral(ch chan<- any, fileName string, structType reflect.Type) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		close(ch)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read the opening bracket of the array
	if _, err := decoder.Token(); err != nil {
		fmt.Println("Error reading opening bracket:", err)
		close(ch)
		return
	}

	// While the array contains values
	for decoder.More() {
		obj := reflect.New(structType).Interface()
		if err := decoder.Decode(obj); err != nil {
			fmt.Println("Error decoding JSON object:", err)
			continue
		}
		ch <- obj
	}

	// Read the closing bracket of the array
	if _, err := decoder.Token(); err != nil {
		fmt.Println("Error reading closing bracket:", err)
	}

	close(ch)
}
