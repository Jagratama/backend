package dto

type UserResponse struct {
	ID         uint   `json:"id"`
	RoleID     uint   `json:"role_id"`
	PositionID uint   `json:"position_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ImagePath  string `json:"image_path"`

	Role     Role     `json:"role"`
	Position Position `json:"position"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Position struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	RequiresSignatures bool   `json:"requires_signatures"`
}

type UserDocumentResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImagePath string `json:"image_path"`
}
