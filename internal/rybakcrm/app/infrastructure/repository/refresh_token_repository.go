package repository

import (
	"context"
	"crm-backend/internal/rybakcrm/app/domain/repository"
	"crm-backend/internal/rybakcrm/app/infrastructure/lib"
	"crm-backend/internal/rybakcrm/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"time"
)

type RefreshTokenRepository struct {
	cfg   *config.Config
	redis *redis.Client
}

func NewRefreshTokenRepository(cfg *config.Config, redis *redis.Client) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		cfg:   cfg,
		redis: redis,
	}
}

func (r *RefreshTokenRepository) GenerateNewToken(accessTokenId string, userId int32) (*jwt.Token, string) {
	refreshTokenId, _ := lib.GenerateTokenId(20)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &repository.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        refreshTokenId,
			ExpiresAt: time.Now().UTC().Add(r.cfg.JWT.RefreshTokenTTL).Unix(),
		},
		AccessTokenId: accessTokenId,
		UserId:        userId,
	})

	return refreshToken, refreshTokenId
}

func (r *RefreshTokenRepository) SaveToken(token *jwt.Token) (string, error) {
	claims, _ := token.Claims.(*repository.RefreshTokenClaims)
	signedToken, err := token.SignedString([]byte(r.cfg.JWT.SignKey))
	if err != nil {
		//todo logger
		return "", err
	}

	r.redis.Set(context.Background(), claims.Id, signedToken, r.cfg.JWT.RefreshTokenTTL)

	return signedToken, nil
}

func (r *RefreshTokenRepository) ParseToken(token string) (*repository.RefreshTokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &repository.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(r.cfg.JWT.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*repository.RefreshTokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *RefreshTokenClaims")
	}

	return claims, nil
}

func (r *RefreshTokenRepository) IsTokenRevoked(id string) (bool, error) {
	result := r.redis.Exists(context.Background(), id)
	return result.Val() == 0, result.Err()
}

func (r *RefreshTokenRepository) RevokeToken(id string) error {
	result := r.redis.Exists(context.Background(), id)
	if result.Val() > 0 {
		result = r.redis.Del(context.Background(), id)
	}

	return result.Err()
}
