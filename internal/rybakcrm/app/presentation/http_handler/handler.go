package http_handler

import (
	"crm-backend/internal/rybakcrm/app/application/interactors"
	"crm-backend/internal/rybakcrm/app/application/usecase"
	"crm-backend/internal/rybakcrm/config"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	cfg    *config.Config
	logger *slog.Logger
	hc     *usecase.Healthcheck
	auth   *interactors.AuthInteractor
	dep    *interactors.DepartmentInteractor
}

func NewHandler(
	cfg *config.Config,
	log *slog.Logger,
	healthCheck *usecase.Healthcheck,
	auth *interactors.AuthInteractor,
	dep *interactors.DepartmentInteractor,
) *Handler {
	return &Handler{
		cfg:    cfg,
		logger: log,
		hc:     healthCheck,
		auth:   auth,
		dep:    dep,
	}
}

func (h *Handler) respondError(c *gin.Context, statusCode int, message interface{}) {
	h.logger.Error("", "message", message)
	c.AbortWithStatusJSON(statusCode, gin.H{
		"message": message,
	})
}

func (h *Handler) respondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
