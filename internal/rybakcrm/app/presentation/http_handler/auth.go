package http_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"crm-backend/internal/rybakcrm/app/application/dto"
)

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) logIn(c *gin.Context) {
	var input loginInput

	ctx := c.Request.Context()

	if err := c.BindJSON(&input); err != nil {
		h.logger.Debug("binding error", "error", err)
		h.respondError(c, http.StatusBadRequest, "invalid data")
		return
	}

	loginDto := &dto.LoginRequestDto{
		Username: input.Username,
		Password: input.Password,
	}

	response, err := h.auth.LogIn(ctx, loginDto)

	if err != nil {
		h.logger.Debug("login error", "error", err)
		h.respondError(c, http.StatusInternalServerError, "smth went wrong")
		return
	}

	http.SetCookie(c.Writer, h.getRefreshTokenCookie(response.RefreshToken, int(h.cfg.JWT.RefreshTokenTTL/time.Second)))

	h.respondSuccess(c, http.StatusOK, gin.H{
		"access_token": response.AccessToken,
		"token_type":   response.TokenType,
		"user":         response.User,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	ctx := c.Request.Context()

	token, err := c.Cookie("RefreshToken")
	if err != nil {
		h.logger.Debug("token error", "error", err)
		h.respondError(c, http.StatusBadRequest, "invalid token")
		return
	}

	request := &dto.RefreshTokenRequestDto{
		RefreshToken: token,
	}

	response, err := h.auth.RefreshToken(ctx, request)

	if err != nil {
		h.logger.Debug("refresh error", "error", err)
		h.respondError(c, http.StatusInternalServerError, "internal server error")
	}

	http.SetCookie(c.Writer, h.getRefreshTokenCookie(response.RefreshToken, int(h.cfg.JWT.RefreshTokenTTL/time.Second)))
	h.respondSuccess(c, http.StatusOK, gin.H{
		"access_token": response.AccessToken,
	})
}

func (h *Handler) logout(c *gin.Context) {
	http.SetCookie(c.Writer, h.getRefreshTokenCookie("deleted", -int(h.cfg.JWT.RefreshTokenTTL/time.Second)))
	h.respondSuccess(c, http.StatusNoContent, "")
}

func (h *Handler) getRefreshTokenCookie(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     "RefreshToken",
		Value:    token,
		MaxAge:   maxAge,
		Path:     "/api/auth",
		Secure:   false,
		HttpOnly: true,
	}
}
