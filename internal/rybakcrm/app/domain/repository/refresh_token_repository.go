package repository

import "github.com/dgrijalva/jwt-go"

type RefreshTokenClaims struct {
	jwt.StandardClaims
	AccessTokenId string `json:"access_token_id"`
	UserId        int32  `json:"user_id"`
}

type RefreshTokenRepository interface {
	GenerateNewToken(accessTokenId string, userId int32) (*jwt.Token, string)
	SaveToken(token *jwt.Token) (string, error)
	ParseToken(token string) (*RefreshTokenClaims, error)
	IsTokenRevoked(id string) (bool, error)
	RevokeToken(id string) error
}
