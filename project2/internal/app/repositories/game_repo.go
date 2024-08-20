package repositories

import (
	"encoding/json"
	"errors"
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
	if _, err := os.Stat(config.GamesFile); err == nil {
		go loadAllGames()
	}
}

type gameRepo struct {
}

func NewGameRepo() interfaces.GameRepository {
	return &gameRepo{}
}

func (g *gameRepo) GetGameByID(gameId string) (*entities.Game, error) {
	game, exists := globals.GamesMap[gameId]
	if !exists {
		return nil, errors.New("game Not Found")
	}
	return &game, nil
}

func (g *gameRepo) GetAllGames() ([]entities.Game, error) {
	var games []entities.Game

	for _, game := range globals.GamesMap {
		games = append(games, game)
	}
	return games, nil
}

func (g *gameRepo) CreateGame(game *entities.Game) error {

	file, err := os.OpenFile(config.GamesFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var games []entities.Game

	// Decode existing games from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&games); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing users:", err)
		return err
	}

	// Append the new game to the games array
	games = append(games, *game)

	// Truncate the file to overwrite it with the updated games array
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
	if err := encoder.Encode(games); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	// add the game to the GamesMap
	globals.GamesMap[game.GameId] = *game
	return nil

}

func (g *gameRepo) DeleteGame(gameId string) error {
	file, err := os.OpenFile(config.GamesFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var games []entities.Game

	// Decode existing games from the file, if the file is not empty
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&games); err != nil && err != io.EOF {
		fmt.Println("Error decoding existing users:", err)
		return err
	}

	// Iterate through the games slice to find the game with the matching GameId
	for i, game := range games {
		if game.GameId == gameId {
			// Remove the game at index i
			games = append(games[:i], games[i+1:]...)
			break
		}
	}

	// Truncate the file to overwrite it with the updated games array
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
	if err := encoder.Encode(games); err != nil {
		fmt.Println("error encoding data to file: %w", err)
		return err
	}

	delete(globals.GamesMap, gameId)
	return nil
}

func loadAllGames() {
	gameDataChan := make(chan any)
	go utils.StreamJSONObjects(gameDataChan, config.GamesFile, reflect.TypeOf(entities.Game{}))

	for game := range gameDataChan {
		gameData, ok := game.(*entities.Game)
		if !ok {
			fmt.Println("Error: received data is not of type entities.Game")
			continue
		}
		globals.GamesMap[gameData.GameId] = *gameData
	}
}
