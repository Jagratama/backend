package dto

type AuthResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	Position     string `json:"position"`
	Image        string `json:"image"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
