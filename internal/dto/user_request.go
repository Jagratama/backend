package dto

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=50"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=50"`
}
