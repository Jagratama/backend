package dto

type UpdatePasswordRequest struct {
	NewPassword     string `json:"new_password" validate:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=50"`
}
