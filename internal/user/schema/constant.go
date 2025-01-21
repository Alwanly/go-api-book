package schema

import (
	"github.com/Alwanly/go-codebase/model"
	"github.com/Alwanly/go-codebase/pkg/middleware"
)

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRegisterResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ProfileRequest struct {
	AuthUserData *middleware.AuthUserData
}

type ProfileResponse struct {
	model.User
}
