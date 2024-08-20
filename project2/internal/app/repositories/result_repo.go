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
	if _, err := os.Stat(config.ResultsFile); err == nil {
		go loadAllResults()
	}
}

type resultRepo struct {
}

func NewResultRepo() interfaces.ResultRepository {
	return &resultRepo{}
}

func (r *resultRepo) AddResult(result *entities.Result) error {
	file, err := os.OpenFile(config.ResultsFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var results []entities.Result

	// Decode existing results from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&results); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing results:", err)
		return err
	}

	// Append the new result to the results array
	results = append(results, *result)

	// Truncate the file to overwrite it with the updated results array
	if err := file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}

	// Move the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return err
	}

	// Encode the updated games array to JSON and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(results); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	// add the result to the resultsMap
	globals.ResultsMap[result.ResultId] = *result
	return nil

}

func (r *resultRepo) RemoveResult(resultId string) error {
	file, err := os.OpenFile(config.ResultsFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var results []entities.Result

	// Decode existing results from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&results); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing results:", err)
		return err
	}

	// remove the result from the results array
	for i, result := range results {
		if resultId == result.ResultId {
			results = append(results[:i], results[i+1:]...)
		}
	}

	// Truncate the file to overwrite it with the updated results array
	if err := file.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}

	// Move the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return err
	}

	// Encode the updated games array to JSON and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: set indentation for pretty printing
	if err := encoder.Encode(results); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	// remove the result from the resultsMap
	delete(globals.ResultsMap, resultId)
	return nil

}

func (r *resultRepo) FindResult(resultId string) *entities.Result {
	result := globals.ResultsMap[resultId]
	return &result
}

func (r *resultRepo) GetAllResults() []*entities.Result {
	var results []*entities.Result

	for _, result := range globals.ResultsMap {
		results = append(results, &result)
	}
	return results
}

func loadAllResults() {
	resultDataChan := make(chan any)
	go utils.StreamJSONObjects(resultDataChan, config.ResultsFile, reflect.TypeOf(entities.Result{}))

	for result := range resultDataChan {
		resultData, ok := result.(*entities.Result)
		if !ok {
			fmt.Println("Error: received data is not of type *entities.Result")
			continue
		}
		globals.ResultsMap[resultData.ResultId] = *resultData
	}
}
