package profile

import "github.com/google/uuid"

type createProfileInput struct {
	Name     string `json:"name" validate:"required,max=32" example:"abhinav"`
	Email    string `json:"email" validate:"required,email,max=32" example:"asrivatsa6@gmail.com"`
	Password string `json:"password" validate:"required,min=8,max=32" example:"something123"`
}

type loginInput struct {
	Email    string `json:"email" validate:"required,email" example:"asrivatsa6@gmail.com"`
	Password string `json:"password" validate:"required" example:"something123"`
}

type profileOutput struct {
	Id        uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name      string    `json:"name" example:"abhinav"`
	Email     string    `json:"email" example:"asrivatsa6@gmail.com"`
	Verified  bool      `json:"verified" example:"false"`
	CreatedAt int64     `json:"created_at" example:"1710579130"`
	UpdatedAt int64     `json:"updated_at" example:"1710579130"`
}

type changePasswordInput struct {
	OldPassword string `json:"old_password" validate:"required" example:"something123"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=32" example:"something1234"`
}

type forgotEmailInput struct {
	Email string `json:"email" validate:"required,email" example:"asrivatsa6@gmail.com"`
}
