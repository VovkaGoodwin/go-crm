package repository

import (
	"context"
	"github.com/dgrijalva/jwt-go"
)

const SignKey = "asd-(#*$;adsl3tto-4551lf9458mbv"

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserId int32 `json:"user_id"`
}

type AccessTokenRepository interface {
	GenerateNewToken(userId int32) (*jwt.Token, string)
	ParseToken(token string) (*AccessTokenClaims, error)
	SaveToken(ctx context.Context, token *jwt.Token) (string, error)
	RevokeToken(ctx context.Context, id string) error
	IsTokenRevoked(ctx context.Context, id string) (bool, error)
}
