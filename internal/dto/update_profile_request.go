package dto

type UpdateProfileRequest struct {
	Name    string `json:"name" validate:"required"`
	ImageID *uint  `json:"image_id" validate:"required"`
}
