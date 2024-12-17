package repository

import "github.com/dgrijalva/jwt-go"

const SignKey = "asd-(#*$;adsl3tto-4551lf9458mbv"

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserId int32 `json:"user_id"`
}

type AccessTokenRepository interface {
	GenerateNewToken(userId int32) (*jwt.Token, string)
	SaveToken(token *jwt.Token) (string, error)
	RevokeToken(id string) error
	ParseToken(token string) (*AccessTokenClaims, error)
	IsTokenRevoked(id string) (bool, error)
}
