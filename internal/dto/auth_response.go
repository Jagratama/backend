package dto

type AuthResponse struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
