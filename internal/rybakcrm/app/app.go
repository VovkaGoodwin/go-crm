package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crm-backend/internal/rybakcrm/config"
)

type App struct {
	log    *slog.Logger
	server *http.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	log := initLog(cfg)
	postgres := initPostgres(cfg)
	redis := initRedis(cfg)

	initContainer(cfg, postgres, redis)

	router := initRouter(ctx, cfg, log)
	server := initHttpServer(cfg, router)

	return &App{
		log,
		server,
	}
}

func (a *App) Start(ctx context.Context) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.log.Error(err.Error())
		}
	}()

	a.log.Info("Server started")
	<-done
	a.log.Info("Stopping server")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Error("server shutdown error", slog.String("error", err.Error()))
		return
	}

	a.log.Info("server stopped")
}
