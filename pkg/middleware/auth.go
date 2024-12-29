package middleware

import (
	"crm-backend/internal/rybakcrm/app/application/interactors"
	"crm-backend/internal/rybakcrm/app/domain/service"
	"crm-backend/pkg/contextutil"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

func Authorize(
	logger *slog.Logger,
	authService *service.AuthService,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader(authorizationHeader)
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			logger.Debug("no authorization header")
			return
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != interactors.TokenType || len(headerParts[1]) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			logger.Debug("invalid authorization header", slog.String("header", header))
			return
		}

		ctx := c.Request.Context()

		userID, err := authService.ParseAccessToken(ctx, headerParts[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			logger.Debug("invalid token", slog.String("token", headerParts[1]), slog.Any("error", err))
			return
		}

		ctx = contextutil.SetCurrentUserID(ctx, userID)
		c.Request.WithContext(ctx)

		c.Next()
	}
}
