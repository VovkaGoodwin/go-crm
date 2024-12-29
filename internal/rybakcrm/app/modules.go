package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"crm-backend/internal/rybakcrm/app/application/interactors"
	"crm-backend/internal/rybakcrm/app/application/usecase"
	"crm-backend/internal/rybakcrm/app/infrastructure/database"
	"crm-backend/internal/rybakcrm/app/presentation/http_handler"
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
	db, err := database.NewPostgresDb(cfg)
	if err != nil {
		panic("database initializing" + err.Error())
	}

	return db
}

func initRedis(cfg *config.Config) *redis.Client {
	return database.NewRedisDb(cfg)
}

func initHandler(
	cfg *config.Config,
	logger *slog.Logger,
) *gin.Engine {
	handler := http_handler.NewHandler(
		cfg,
		logger,
		usecase.NewHealthCheckUseCase(),
		interactors.NewAuthInteractor(cfg, logger, AuthService()),
		interactors.NewDepartmentInteractor(cfg, logger, DepartmentService()),
	)

	return handler.InitRoutes(AuthService())
}
