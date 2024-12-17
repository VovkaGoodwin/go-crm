package interactors

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"crm-backend/internal/rybakcrm/app/application/http_response"
	"crm-backend/internal/rybakcrm/app/domain/service"
	"crm-backend/internal/rybakcrm/config"
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

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthInteractor) LogIn(ctx *gin.Context) {
	var input loginInput

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, http_response.Error{
			http.StatusBadRequest,
			"invalid data",
		})
		a.log.Error("Input binding error", slog.String("error", err.Error()))
		return
	}

	accessToken, refreshToken, user, err := a.authService.Login(input.Username, input.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Wrong user credentials",
		})
		a.log.Error("LogIn error", slog.String("error", err.Error()))
		return
	}

	http.SetCookie(ctx.Writer, a.getRefreshTokenCookie(refreshToken, int(a.cfg.JWT.RefreshTokenTTL/time.Second)))

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"token_type":   TokenType,
		"user":         user,
	})
}

func (a *AuthInteractor) RefreshToken(ctx *gin.Context) {
	token, err := ctx.Cookie("RefreshToken")
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		a.log.Error("Refresh token nod found")
		return
	}

	newAccessToken, newRefreshToken, err := a.authService.RefreshToken(token)
	if err != nil {
		//h.abortWithMessage(c, http.StatusInternalServerError, fmt.Sprintf("token refreshing error: %s", err.Error()))
		return
	}

	http.SetCookie(ctx.Writer, a.getRefreshTokenCookie(newRefreshToken, int(a.cfg.JWT.RefreshTokenTTL/time.Second)))

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"access_token": newAccessToken,
	})
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
