package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"project2/internal/domain/entities"
	interfaces "project2/internal/domain/interfaces/repository"
)

type gameRepo struct {
	db *sql.DB
}

func NewGameRepo(db *sql.DB) interfaces.GameRepository {
	return &gameRepo{
		db: db,
	}
}

// FetchGameByID retrieves a game by its ID.
func (r *gameRepo) FetchGameByID(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	query := `SELECT game_id, game_name, min_players, max_players, instances, is_active, created_at, updated_at FROM games WHERE game_id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var game entities.Game
	err := row.Scan(&game.GameID, &game.GameName, &game.MinPlayers, &game.MaxPlayers, &game.Instances, &game.IsActive, &game.CreatedAt, &game.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No game found
		}
		return nil, fmt.Errorf("failed to fetch game by ID: %w", err)
	}

	return &game, nil
}

// FetchAllGames retrieves all games from the database.
func (r *gameRepo) FetchAllGames(ctx context.Context) ([]entities.Game, error) {
	query := `SELECT game_id, game_name, min_players, max_players, instances, is_active, created_at, updated_at FROM games`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all games: %w", err)
	}
	defer rows.Close()

	var games []entities.Game
	for rows.Next() {
		var game entities.Game
		if err := rows.Scan(&game.GameID, &game.GameName, &game.MinPlayers, &game.MaxPlayers, &game.Instances, &game.IsActive, &game.CreatedAt, &game.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan game row: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over games: %w", err)
	}

	return games, nil
}

// CreateGame inserts a new game into the database and returns the created game ID.
func (r *gameRepo) CreateGame(ctx context.Context, game *entities.Game) (uuid.UUID, error) {
	query := `INSERT INTO games (game_name, min_players, max_players, instances, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING game_id`
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, query, game.GameName, game.MinPlayers, game.MaxPlayers, game.Instances, game.IsActive).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create game: %w", err)
	}
	return id, nil
}

// DeleteGame removes a game from the database by its ID.
func (r *gameRepo) DeleteGame(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM games WHERE game_id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete game: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no game found with ID %s", id)
	}

	return nil
}

// UpdateGameStatus updates the status of a game
func (r *gameRepo) UpdateGameStatus(ctx context.Context, gameID uuid.UUID, status bool) error {
	query := `UPDATE games SET is_active = $1, updated_at = CURRENT_TIMESTAMP WHERE game_id = $2`

	// Execute the update query
	_, err := r.db.ExecContext(ctx, query, status, gameID)
	if err != nil {
		return fmt.Errorf("failed to update game status: %w", err)
	}
	return nil
}

//func (g *gameRepo) GetGameByID(gameId primitive.ObjectID) (*entities.Game, error) {
//	var game entities.Game
//	filter := bson.D{{"_id", gameId}}
//	err := g.collection.FindOne(context.Background(), filter).Decode(&game)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			return nil, fmt.Errorf("game not found")
//		}
//		fmt.Println("Error finding game:", err)
//		return nil, err
//	}
//	return &game, nil
//}
//
//func (g *gameRepo) GetAllGames() ([]entities.Game, error) {
//	var games []entities.Game
//	cursor, err := g.collection.Find(context.Background(), bson.D{})
//	if err != nil {
//		fmt.Println("Error finding games:", err)
//		return nil, err
//	}
//	defer cursor.Close(context.Background())
//
//	for cursor.Next(context.Background()) {
//		var game entities.Game
//		if err := cursor.Decode(&game); err != nil {
//			fmt.Println("Error decoding game:", err)
//			return nil, err
//		}
//		games = append(games, game)
//	}
//	if err := cursor.Err(); err != nil {
//		fmt.Println("Cursor error:", err)
//		return nil, err
//	}
//
//	return games, nil
//}
//
//func (g *gameRepo) CreateGame(game *entities.Game) error {
//	_, err := g.collection.InsertOne(context.Background(), game)
//	if err != nil {
//		fmt.Println("Error inserting game:", err)
//		return err
//	}
//	return nil
//}
//
//func (g *gameRepo) DeleteGame(gameId primitive.ObjectID) error {
//	_, err := g.collection.DeleteOne(context.Background(), bson.D{{"_id", gameId}})
//	if err != nil {
//		fmt.Println("Error deleting game:", err)
//		return err
//	}
//	return nil
//}
