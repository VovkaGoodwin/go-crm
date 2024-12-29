package interactors

import (
	"context"
	"crm-backend/internal/rybakcrm/app/application/dto"
	"crm-backend/internal/rybakcrm/app/domain/service"
	"crm-backend/internal/rybakcrm/config"
	"log/slog"
	"net/http"
)

const TokenType = "Bearer"

type AuthInteractor struct {
	cfg         *config.Config
	log         *slog.Logger
	authService *service.AuthService
}

func NewAuthInteractor(
	cfg *config.Config,
	log *slog.Logger,
	authService *service.AuthService,
) *AuthInteractor {
	return &AuthInteractor{
		cfg:         cfg,
		log:         log,
		authService: authService,
	}
}

func (a *AuthInteractor) LogIn(ctx context.Context, request *dto.LoginRequestDto) (*dto.LoginResponseDto, error) {

	accessToken, refreshToken, user, err := a.authService.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
		TokenType:    TokenType,
	}, nil
}

func (a *AuthInteractor) RefreshToken(ctx context.Context, request *dto.RefreshTokenRequestDto) (*dto.RefreshTokenResponseDto, error) {
	accessToken, refreshToken, err := a.authService.RefreshToken(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	response := &dto.RefreshTokenResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (a *AuthInteractor) Logout(ctx context.Context, token string) {

}

func (a *AuthInteractor) getRefreshTokenCookie(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     "RefreshToken",
		Value:    token,
		MaxAge:   maxAge,
		Path:     "/api/auth",
		Secure:   false,
		HttpOnly: true,
	}
}
