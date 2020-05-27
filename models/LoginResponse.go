package models

// LoginResponse handles JWT-Response-structure
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
