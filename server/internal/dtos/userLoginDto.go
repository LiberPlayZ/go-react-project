package dtos

// dto represents a  login user in the system
type UserLoginDto struct {
	Email    string `json:"email"` 
	Password string `json:"password"`
}
