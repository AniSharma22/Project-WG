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

type notificationRepo struct {
	db *sql.DB
}

func NewNotificationRepo(db *sql.DB) interfaces.NotificationRepository {
	return &notificationRepo{db: db}
}

// CreateNotification inserts a new notification into the database and returns the created notification ID.
func (r *notificationRepo) CreateNotification(ctx context.Context, notification *entities.Notification) (uuid.UUID, error) {
	query := `INSERT INTO notifications (user_id, message, is_read) VALUES ($1, $2, $3) RETURNING notification_id`
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, query, notification.UserID, notification.Message, notification.IsRead).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create notification: %w", err)
	}
	return id, nil
}

// FetchNotificationByID retrieves a notification by its ID.
func (r *notificationRepo) FetchNotificationByID(ctx context.Context, id uuid.UUID) (*entities.Notification, error) {
	query := `SELECT notification_id, user_id, message, is_read, created_at FROM notifications WHERE notification_id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var notification entities.Notification
	err := row.Scan(&notification.NotificationID, &notification.UserID, &notification.Message, &notification.IsRead, &notification.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No notification found
		}
		return nil, fmt.Errorf("failed to fetch notification by ID: %w", err)
	}

	return &notification, nil
}

// FetchUserNotifications retrieves all notifications for a specific user.
func (r *notificationRepo) FetchUserNotifications(ctx context.Context, userID uuid.UUID) ([]entities.Notification, error) {
	query := `SELECT notification_id, user_id, message, is_read, created_at FROM notifications WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user notifications: %w", err)
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		if err := rows.Scan(&notification.NotificationID, &notification.UserID, &notification.Message, &notification.IsRead, &notification.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan notification row: %w", err)
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over notifications: %w", err)
	}

	return notifications, nil
}

// MarkNotificationAsRead marks a notification as read by updating its is_read status.
func (r *notificationRepo) MarkNotificationAsRead(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE notifications SET is_read = TRUE WHERE notification_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to mark notification as read: %w", err)
	}
	return nil
}

// DeleteNotificationByID deletes a notification from the database by its ID.
func (r *notificationRepo) DeleteNotificationByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM notifications WHERE notification_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}
