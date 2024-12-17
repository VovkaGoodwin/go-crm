package app

import (
	"context"
	"crm-backend/internal/rybakcrm/app/infrastructure/database"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"crm-backend/internal/rybakcrm/app/application/interactors"
	"crm-backend/internal/rybakcrm/app/application/usecase"
	"crm-backend/internal/rybakcrm/config"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func initLog(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func initHttpServer(cfg *config.Config, handler http.Handler) *http.Server {

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		Handler:      handler,
	}

	return server
}

func initPostgres(cfg *config.Config) *sqlx.DB {
	db, _ := database.NewPostgresDb(cfg)

	return db
}

func initRedis(cfg *config.Config) *redis.Client {
	return database.NewRedisDb(cfg)
}

func initRouter(
	ctx context.Context,
	cfg *config.Config,
	log *slog.Logger,
) *gin.Engine {
	router := gin.New()

	router.Use(func(ctx *gin.Context) {
		log.Info("request received", slog.String("url", ctx.Request.URL.String()))
	}, func(c *gin.Context) {
		c.Set("ctx", ctx)
	})

	api := router.Group("/api")
	{
		api.GET("/healthcheck", usecase.NewHealthCheckUseCase().Handle)

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", interactors.NewAuthInteractor(cfg, log, AuthService()).LogIn)
		}
	}

	return router
}
