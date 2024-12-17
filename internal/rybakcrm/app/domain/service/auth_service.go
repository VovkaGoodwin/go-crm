package service

import (
	"errors"

	"crm-backend/internal/rybakcrm/app/domain/models"
	"crm-backend/internal/rybakcrm/app/domain/repository"
)

type AuthService struct {
	userRepository         repository.UserRepository
	accessTokenRepository  repository.AccessTokenRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewAuthService(
	userRepository repository.UserRepository,
	accessTokenRepository repository.AccessTokenRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		accessTokenRepository:  accessTokenRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (a *AuthService) Login(username, password string) (string, string, models.User, error) {
	user, err := a.userRepository.GetUserByCredentials(username, password)

	if err != nil {
		return "", "", user, err
	}

	accessToken, accessTokenId := a.accessTokenRepository.GenerateNewToken(user.ID)

	refreshToken, _ := a.refreshTokenRepository.GenerateNewToken(accessTokenId, user.ID)

	signedAccessToken, err := a.accessTokenRepository.SaveToken(accessToken)

	if err != nil {
		return "", "", user, err
	}

	signedRefreshToken, err := a.refreshTokenRepository.SaveToken(refreshToken)
	if err != nil {
		return "", "", user, err
	}

	return signedAccessToken, signedRefreshToken, user, nil
}

func (a *AuthService) RefreshToken(token string) (newAccessToken, newRefreshToken string, err error) {

	claims, err := a.refreshTokenRepository.ParseToken(token)
	if err != nil {
		return "", "", err
	}

	isTokenRevoked, err := a.refreshTokenRepository.IsTokenRevoked(claims.Id)
	if err != nil {
		return "", "", err
	}

	if isTokenRevoked {
		return "", "", errors.New("invalid refresh token")
	}

	err = a.accessTokenRepository.RevokeToken(claims.AccessTokenId)
	if err != nil {
		return "", "", err
	}

	err = a.refreshTokenRepository.RevokeToken(claims.Id)
	if err != nil {
		return "", "", err
	}

	accessToken, accessTokenId := a.accessTokenRepository.GenerateNewToken(claims.UserId)
	refreshToken, _ := a.refreshTokenRepository.GenerateNewToken(accessTokenId, claims.UserId)

	newAccessToken, err = a.accessTokenRepository.SaveToken(accessToken)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err = a.refreshTokenRepository.SaveToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (a *AuthService) ParseAccessToken(token string) (int32, error) {

	claims, err := a.accessTokenRepository.ParseToken(token)
	if err != nil {
		return 0, err
	}

	revoked, err := a.accessTokenRepository.IsTokenRevoked(claims.Id)
	if err != nil {
		return 0, err
	}

	if revoked {
		return 0, errors.New("revoked token")
	}

	return claims.UserId, nil
}
