package dto

type UpdateProfileRequest struct {
	Name      string `json:"name" validate:"required"`
	ImagePath string `json:"image_path" validate:"required"`
}
