package dto

import (
	"crm-backend/internal/rybakcrm/app/domain/models"
)

type LoginRequestDto struct {
	Username string
	Password string
}

type LoginResponseDto struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	User         models.User
}

type RefreshTokenRequestDto struct {
	RefreshToken string
}

type RefreshTokenResponseDto struct {
	RefreshToken string
	AccessToken  string
}
