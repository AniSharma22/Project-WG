package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"project2/internal/domain/entities"
	interfaces "project2/internal/domain/interfaces/repository"
	"project2/internal/models"
	"time"
)

type bookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) interfaces.BookingRepository {
	return &bookingRepo{
		db: db,
	}
}

// CreateBooking inserts a new booking into the database and returns the created booking ID.
func (r *bookingRepo) CreateBooking(ctx context.Context, booking *entities.Booking) (uuid.UUID, error) {
	query := `INSERT INTO bookings (slot_id, user_id) VALUES ($1, $2) RETURNING booking_id`
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, query, booking.SlotID, booking.UserID).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create booking: %w", err)
	}
	return id, nil
}

// FetchBookingByID retrieves a booking by its ID.
func (r *bookingRepo) FetchBookingByID(ctx context.Context, id uuid.UUID) (*entities.Booking, error) {
	query := `SELECT booking_id, slot_id, user_id, created_at FROM bookings WHERE booking_id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var booking entities.Booking
	err := row.Scan(&booking.BookingID, &booking.SlotID, &booking.UserID, &booking.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No booking found
		}
		return nil, fmt.Errorf("failed to fetch booking by ID: %w", err)
	}

	return &booking, nil
}

// DeleteBookingByID removes a booking from the database by its ID.
func (r *bookingRepo) DeleteBookingByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM bookings WHERE booking_id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no booking found with ID %s", id)
	}

	return nil
}

// FetchBookingsByUserID retrieves all bookings associated with a specific user ID.
func (r *bookingRepo) FetchBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Booking, error) {
	query := `SELECT booking_id, slot_id, user_id, created_at FROM bookings WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bookings by user ID: %w", err)
	}
	defer rows.Close()

	var bookings []entities.Booking
	for rows.Next() {
		var booking entities.Booking
		if err := rows.Scan(&booking.BookingID, &booking.SlotID, &booking.UserID, &booking.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan booking row: %w", err)
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over bookings: %w", err)
	}

	return bookings, nil
}

// FetchBookingsBySlotID retrieves all bookings for a specific slot ID.
func (r *bookingRepo) FetchBookingsBySlotID(ctx context.Context, slotID uuid.UUID) ([]entities.Booking, error) {
	// Define the query to fetch bookings by slot ID
	query := `SELECT booking_id, slot_id, user_id, created_at 
	          FROM bookings 
	          WHERE slot_id = $1`

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, slotID)
	if err != nil {
		return nil, fmt.Errorf("failed to query bookings by slot ID: %w", err)
	}
	defer rows.Close()

	// Initialize a slice to hold the results
	var bookings []entities.Booking

	// Iterate over the rows
	for rows.Next() {
		var booking entities.Booking
		err := rows.Scan(&booking.BookingID, &booking.SlotID, &booking.UserID, &booking.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}
		bookings = append(bookings, booking)
	}

	// Check for errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bookings, nil
}

// FetchUpcomingBookingsByUserID retrieves all upcoming bookings with game information for a specific user.
func (r *bookingRepo) FetchUpcomingBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	// SQL query to join bookings, slots, and games tables and filter by user ID and future slot start time
	query := `
		SELECT 
			g.game_name, 
			s.slot_date AS date, 
			s.start_time AS start_time, 
			s.end_time AS end_time, 
			ARRAY_AGG(u.username) AS booked_users
		FROM bookings b
		JOIN slots s ON b.slot_id = s.slot_id
		JOIN games g ON s.game_id = g.game_id
		JOIN users u ON b.user_id = u.user_id
		WHERE b.user_id = $1
		  AND s.start_time > NOW()
		GROUP BY g.game_name, s.slot_date, s.start_time, s.end_time
		ORDER BY s.start_time ASC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch upcoming bookings by user ID: %w", err)
	}
	defer rows.Close()

	var bookings []models.Bookings
	for rows.Next() {
		var booking models.Bookings
		var bookedUsers pq.StringArray // PostgresSQL-specific array type

		err := rows.Scan(&booking.GameName, &booking.Date, &booking.StartTime, &booking.EndTime, &bookedUsers)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}

		// Convert pq.StringArray to []string
		booking.BookedUsers = []string(bookedUsers)

		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bookings, nil
}

// UpdateBookingResult updates the result (win/loss) of the specified booking
func (r *bookingRepo) UpdateBookingResult(ctx context.Context, bookingId uuid.UUID, result string) error {
	// Define the SQL query to update the result
	query := `
		UPDATE bookings
		SET result = $1
		WHERE booking_id = $2
	`

	// Execute the query
	_, err := r.db.ExecContext(ctx, query, result, bookingId)
	if err != nil {
		return fmt.Errorf("failed to update booking result: %w", err)
	}

	return nil
}

func (r *bookingRepo) FetchBookingsToUpdateResult(ctx context.Context, userID uuid.UUID) ([]models.Bookings, error) {
	// Define the SQL query
	query := `
		SELECT
			g.game_name,
			s.slot_date,
			s.start_time,
			s.end_time,
			ARRAY_AGG(u.username) AS booked_users
		FROM
			bookings b
			JOIN slots s ON b.slot_id = s.slot_id
			JOIN games g ON s.game_id = g.game_id
			JOIN users u ON b.user_id = u.user_id
		WHERE
			b.user_id = $1
			AND s.end_time < $2
		GROUP BY
			g.game_name, s.slot_date, s.start_time, s.end_time
		ORDER BY
			s.end_time DESC;
	`

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, userID, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch past bookings: %w", err)
	}
	defer rows.Close()

	// Prepare to collect the results
	var bookings []models.Bookings

	// Iterate over the rows
	for rows.Next() {
		var booking models.Bookings
		var bookedUsers []sql.NullString

		// Scan the row into the booking structure
		err := rows.Scan(
			&booking.GameName,
			&booking.Date,
			&booking.StartTime,
			&booking.EndTime,
			pq.Array(&bookedUsers),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking row: %w", err)
		}

		// Convert SQL NullString slice to []string
		for _, user := range bookedUsers {
			if user.Valid {
				booking.BookedUsers = append(booking.BookedUsers, user.String)
			}
		}

		// Append the booking to the result slice
		bookings = append(bookings, booking)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bookings, nil
}

// FetchSlotBookedUsers returns a slice of usernames of all users who have booked the given slot
func (r *bookingRepo) FetchSlotBookedUsers(ctx context.Context, slotId uuid.UUID) ([]string, error) {
	// Define the SQL query to fetch usernames
	query := `
		SELECT u.username
		FROM bookings b
		INNER JOIN users u ON b.user_id = u.user_id
		WHERE b.slot_id = $1
	`

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, slotId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch booked users: %w", err)
	}
	defer rows.Close()

	// Slice to hold the usernames
	var usernames []string

	// Iterate over the rows and append each username to the slice
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, fmt.Errorf("failed to scan username: %w", err)
		}
		usernames = append(usernames, username)
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return usernames, nil
}

func (r *bookingRepo) FetchBookingBySlotAndUserId(ctx context.Context, slotId uuid.UUID, userID uuid.UUID) (models.Bookings, error) {
	query := `
		SELECT
			b.booking_id,
			g.game_name,
			s.slot_date,
			s.start_time,
			s.end_time,
			ARRAY_AGG(u.username) AS booked_users
		FROM
			bookings b
			JOIN slots s ON b.slot_id = s.slot_id
			JOIN games g ON s.game_id = g.game_id
			LEFT JOIN users u ON b.user_id = u.user_id
		WHERE
			b.slot_id = $1
			AND b.user_id = $2
		GROUP BY
			b.booking_id, g.game_name, s.slot_date, s.start_time, s.end_time
	`

	row := r.db.QueryRowContext(ctx, query, slotId, userID)

	var booking models.Bookings
	var bookedUsers []sql.NullString

	err := row.Scan(
		&booking.BookingId,
		&booking.GameName,
		&booking.Date,
		&booking.StartTime,
		&booking.EndTime,
		pq.Array(&bookedUsers),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Bookings{}, nil // No booking found
		}
		return models.Bookings{}, fmt.Errorf("failed to fetch booking by slot and user ID: %w", err)
	}

	// Convert SQL NullString slice to []string for booked users
	for _, user := range bookedUsers {
		if user.Valid {
			booking.BookedUsers = append(booking.BookedUsers, user.String)
		}
	}

	return booking, nil
}
