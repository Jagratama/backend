package dto

type UserResponse struct {
	ID           uint   `json:"id"`
	ImageID      uint   `json:"image_id"`
	RoleID       uint   `json:"role_id"`
	PositionID   uint   `json:"position_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Image        string `json:"image"`
	Organization string `json:"organization"`

	Role     Role     `json:"role"`
	Position Position `json:"position"`
}

type UserPaginationResponse struct {
	Data      []*UserResponse `json:"data"`
	TotalData int             `json:"total_data"`
	Limit     int             `json:"limit"`
	Page      int             `json:"page"`
	TotalPage int             `json:"total_page"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Position struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserImage struct {
	FilePath string `json:"path"`
}

type UserDocumentResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}
