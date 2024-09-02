package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	MobileNumber string    `json:"mobile_number,omitempty" db:"mobile_number"`
	Gender       string    `json:"gender" db:"gender"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
