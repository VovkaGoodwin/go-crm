package repository

import (
	"context"
	"github.com/dgrijalva/jwt-go"
)

type RefreshTokenClaims struct {
	jwt.StandardClaims
	AccessTokenId string `json:"access_token_id"`
	UserId        int32  `json:"user_id"`
}

type RefreshTokenRepository interface {
	GenerateNewToken(accessTokenId string, userId int32) (*jwt.Token, string)
	ParseToken(token string) (*RefreshTokenClaims, error)
	SaveToken(ctx context.Context, token *jwt.Token) (string, error)
	IsTokenRevoked(ctx context.Context, id string) (bool, error)
	RevokeToken(ctx context.Context, id string) error
}
