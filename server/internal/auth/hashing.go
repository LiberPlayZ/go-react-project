package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Hashing struct {
	Cost int
}

func NewHashing(cost int) *Hashing {
	return &Hashing{Cost: cost}
}

// HashPassword generates a bcrypt hash for the given password.
func (p *Hashing) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), p.Cost)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func (p *Hashing) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
