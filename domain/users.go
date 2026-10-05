package domain

import "errors"

// ErrUserNotFound is returned when no user matches the given id.
var ErrUserNotFound = errors.New("user not found")

// User is the entity representing a user in the system.
type User struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
