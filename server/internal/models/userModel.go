package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `json:"id"`       // UUID as the primary key
	Username  string    `json:"username"` // Unique username
	Email     string    `json:"email"`
	Password  string    `json:"password"`   // Hashed password
	Role      string    `json:"role"`       // User role
	CreatedAt time.Time `json:"created_at"` // Timestamp when created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when updated
}
