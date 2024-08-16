package utils

import (
	"github.com/google/uuid"
)

// GetUuid generates a new random UUID and returns it as a string
func GetUuid() (string, error) {
	// Generate a new UUID
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
