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
)

type invitationRepo struct {
	db *sql.DB
}

func NewInvitationRepo(db *sql.DB) interfaces.InvitationRepository {
	return &invitationRepo{db: db}
}

// CreateInvitation inserts a new invitation into the database and returns the created invitation ID.
func (r *invitationRepo) CreateInvitation(ctx context.Context, invitation *entities.Invitation) (uuid.UUID, error) {
	query := `INSERT INTO invitations (inviting_user_id, invited_user_id, slot_id) VALUES ($1, $2, $3) RETURNING invitation_id`
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, query, invitation.InvitingUserID, invitation.InvitedUserID, invitation.SlotID).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create invitation: %w", err)
	}
	return id, nil
}

// DeleteInvitationByID removes an invitation from the database by its ID.
func (r *invitationRepo) DeleteInvitationByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM invitations WHERE invitation_id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete invitation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no invitation found with ID %s", id)
	}

	return nil
}

// UpdateInvitationStatus updates the status of an invitation by its ID.
func (r *invitationRepo) UpdateInvitationStatus(ctx context.Context, id uuid.UUID, status string) error {
	query := `UPDATE invitations SET status = $1 WHERE invitation_id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update invitation status: %w", err)
	}
	return nil
}

// FetchInvitationByID retrieves an invitation by its ID.
func (r *invitationRepo) FetchInvitationByID(ctx context.Context, id uuid.UUID) (*entities.Invitation, error) {
	query := `SELECT invitation_id, inviting_user_id, invited_user_id, status, created_at FROM invitations WHERE invitation_id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var invitation entities.Invitation
	err := row.Scan(&invitation.InvitationID, &invitation.InvitingUserID, &invitation.InvitedUserID, &invitation.Status, &invitation.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No invitation found
		}
		return nil, fmt.Errorf("failed to fetch invitation by ID: %w", err)
	}

	return &invitation, nil
}

// FetchUserInvitations retrieves all invitations sent to or from a specific user.
func (r *invitationRepo) FetchUserInvitations(ctx context.Context, userID uuid.UUID) ([]entities.Invitation, error) {
	query := `SELECT invitation_id, inviting_user_id, invited_user_id, status, created_at FROM invitations WHERE inviting_user_id = $1 OR invited_user_id = $2`
	rows, err := r.db.QueryContext(ctx, query, userID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user invitations: %w", err)
	}
	defer rows.Close()

	var invitations []entities.Invitation
	for rows.Next() {
		var invitation entities.Invitation
		if err := rows.Scan(&invitation.InvitationID, &invitation.InvitingUserID, &invitation.InvitedUserID, &invitation.Status, &invitation.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan invitation row: %w", err)
		}
		invitations = append(invitations, invitation)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over invitations: %w", err)
	}

	return invitations, nil
}

func (r *invitationRepo) FetchUserPendingInvitations(ctx context.Context, userID uuid.UUID) ([]models.Invitations, error) {
	query := `
		SELECT
			i.invitation_id,
			s.slot_id,
			g.game_name,
			s.slot_date,
			s.start_time,
			s.end_time,
			ARRAY_AGG(u.username) AS booked_users,
			inviter.username AS invited_by_username
		FROM
			invitations i
			JOIN slots s ON i.slot_id = s.slot_id
			JOIN games g ON s.game_id = g.game_id
			LEFT JOIN bookings b ON s.slot_id = b.slot_id
			LEFT JOIN users u ON b.user_id = u.user_id
			JOIN users inviter ON i.inviting_user_id = inviter.user_id
		WHERE
			i.invited_user_id = $1
			AND i.status = 'pending'
		GROUP BY
			i.invitation_id, s.slot_id, g.game_name, s.slot_date, s.start_time, s.end_time, inviter.username
		ORDER BY
			s.start_time ASC;
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending invitations: %w", err)
	}
	defer rows.Close()

	var invitations []models.Invitations
	for rows.Next() {
		var invitation models.Invitations
		var bookedUsers []sql.NullString

		err := rows.Scan(
			&invitation.InvitationId,
			&invitation.SlotId,
			&invitation.GameName,
			&invitation.Date,
			&invitation.StartTime,
			&invitation.EndTime,
			pq.Array(&bookedUsers),
			&invitation.InvitedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invitation: %w", err)
		}

		// Convert SQL NullString slice to []string
		for _, user := range bookedUsers {
			if user.Valid {
				invitation.BookedUsers = append(invitation.BookedUsers, user.String)
			}
		}

		invitations = append(invitations, invitation)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return invitations, nil
}

func (r *invitationRepo) FetchInvitationByUserAndSlot(ctx context.Context, invitingUserID uuid.UUID, invitedUserID uuid.UUID, slotID uuid.UUID) (*entities.Invitation, error) {
	query := `
		SELECT id, inviting_user_id, invited_user_id, slot_id
		FROM invitations
		WHERE inviting_user_id = $1 AND invited_user_id = $2 AND slot_id = $3
	`

	row := r.db.QueryRowContext(ctx, query, invitingUserID, invitedUserID, slotID)

	var invitation entities.Invitation
	err := row.Scan(&invitation.InvitationID, &invitation.InvitingUserID, &invitation.InvitedUserID, &invitation.SlotID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No matching invitation found
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch invitation by user and slot: %w", err)
	}

	return &invitation, nil
}
