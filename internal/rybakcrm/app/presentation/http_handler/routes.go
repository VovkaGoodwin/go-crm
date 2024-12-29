package http_handler

import (
	"crm-backend/internal/rybakcrm/app/domain/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *Handler) InitRoutes(service *service.AuthService) *gin.Engine {
	router := gin.New()

	router.Use(func(ctx *gin.Context) {
		h.logger.Info("request received",
			slog.String("url", ctx.Request.URL.String()),
			slog.String("method", ctx.Request.Method),
		)
	})

	api := router.Group("/api")
	{
		api.GET("/healthcheck", h.healthCheck)

		auth := api.Group("/auth")
		{
			auth.POST("/login", h.logIn)
			auth.GET("/refresh", h.refresh)
			auth.GET("/logout", h.logout)
		}

		withAuthorize := api.Group("", func(c *gin.Context) { c.Next() })
		{
			department := withAuthorize.Group("/departments")
			{
				department.GET("", h.getAllDepartments)
				department.POST("", h.createDepartment)
				department.GET("/:id", h.getDepartment)
				department.PUT("/:id", h.updateDepartment)
				department.DELETE("/:id", h.deleteDepartment)
			}

			user := withAuthorize.Group("/user")
			{
				user.POST("")
			}
		}
	}

	return router
}
