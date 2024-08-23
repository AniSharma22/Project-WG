package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func WriteToFile(filename string, inputMap interface{}, structType reflect.Type) {

	// Ensure inputMap is a map
	mapValue := reflect.ValueOf(inputMap)
	if mapValue.Kind() != reflect.Map {
		fmt.Println("input is not a map")
		return
	}

	// Open or create the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating file: %w", err)
		return
	}
	defer file.Close()

	// Convert map to slice of the specified struct type
	var slice []interface{}
	for _, key := range mapValue.MapKeys() {
		keyValue := mapValue.MapIndex(key)

		// Create a new instance of the struct type
		structInstance := reflect.New(structType).Elem()

		// Dynamically set the struct fields
		for i := 0; i < structType.NumField(); i++ {
			field := structType.Field(i)
			fieldValue := structInstance.Field(i)

			if fieldValue.CanSet() {
				switch i {
				case 0: // Assume first field is for the key
					if fieldValue.Kind() == reflect.String {
						fieldValue.SetString(key.String())
					} else {
						fieldValue.Set(key)
					}
				case 1: // Assume second field is for the value
					if fieldValue.Kind() == reflect.String {
						fieldValue.SetString(keyValue.String())
					} else {
						fieldValue.Set(keyValue)
					}
				default:
					// For any additional fields, set them to zero value
					fieldValue.Set(reflect.Zero(field.Type))
				}
			}
		}

		// Append the struct instance to the slice
		slice = append(slice, structInstance.Interface())
	}

	// Create a JSON encoder and write the slice to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(slice); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return
	}

}
