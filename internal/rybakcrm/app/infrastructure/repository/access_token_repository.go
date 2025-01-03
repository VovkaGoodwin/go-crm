package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"

	"crm-backend/internal/rybakcrm/app/domain/repository"
	"crm-backend/internal/rybakcrm/config"
	jwtutils "crm-backend/pkg/jwt"
)

type AccessTokenRepository struct {
	cfg   *config.Config
	redis *redis.Client
}

func NewAccessTokenRepository(cfg *config.Config, redis *redis.Client) *AccessTokenRepository {
	return &AccessTokenRepository{
		cfg:   cfg,
		redis: redis,
	}
}

func (a *AccessTokenRepository) GenerateNewToken(userId int32) (*jwt.Token, string) {
	accessTokenId, _ := jwtutils.GenerateTokenId(20)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &repository.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        accessTokenId,
			ExpiresAt: time.Now().UTC().Add(a.cfg.JWT.AccessTokenTTL).Unix(),
		},
		UserId: userId,
	})

	return accessToken, accessTokenId
}

func (a *AccessTokenRepository) SaveToken(ctx context.Context, token *jwt.Token) (string, error) {
	claims, _ := token.Claims.(*repository.AccessTokenClaims)
	signedToken, err := token.SignedString([]byte(a.cfg.JWT.SignKey))
	if err != nil {
		//todo logger
		return "", err
	}

	a.redis.Set(ctx, claims.Id, signedToken, a.cfg.JWT.AccessTokenTTL)

	return signedToken, nil
}

func (a *AccessTokenRepository) RevokeToken(ctx context.Context, id string) error {
	result := a.redis.Exists(ctx, id)
	if result.Val() > 0 {
		result = a.redis.Del(ctx, id)
	}

	return result.Err()
}

func (a *AccessTokenRepository) ParseToken(token string) (*repository.AccessTokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &repository.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(a.cfg.JWT.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*repository.AccessTokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *AccessTokenClaims")
	}

	return claims, nil
}

func (a *AccessTokenRepository) IsTokenRevoked(ctx context.Context, id string) (bool, error) {
	result := a.redis.Exists(ctx, id)
	return result.Val() == 0, result.Err()
}
