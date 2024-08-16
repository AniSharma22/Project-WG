package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func LoadObjectsFromFile(ch chan<- any, fileName string, structType reflect.Type) {
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
