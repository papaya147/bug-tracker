package profile

import "github.com/google/uuid"

type createProfileRequest struct {
	Name     string `json:"name" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,email,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type profileResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
	Token     string    `json:"token,omitempty"`
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=32"`
}
