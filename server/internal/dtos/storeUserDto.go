package dtos

import (
	"time"

	"github.com/google/uuid"
)

// dto represents a  login user in the system
type StoreUserDto struct {
	ID        uuid.UUID   `json:"id"`       // UUID as the primary key
	Username  string      `json:"username"` // Unique username
	Email     string      `json:"email"`
	Role      string      `json:"role"` // User role
	Todos     []uuid.UUID `json:"todos"`
	CreatedAt time.Time   `json:"created_at"` // Timestamp when created
	UpdatedAt time.Time   `json:"updated_at"` // Timestamp when updated
}
