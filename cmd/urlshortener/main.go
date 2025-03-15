package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"restapi/internal/config"
	"restapi/internal/handlers/url/redirect"
	"restapi/internal/handlers/url/save"
	"restapi/internal/storage/postgresql"

	"github.com/go-chi/chi"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.InitConfig()
	log := setUpLogger(cfg.Env)
	log.Info("start application", slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Address))
	storage, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage")
		return
	}
	router := chi.NewRouter()
	router.Post("/url", save.New(log, storage))
	router.Get("/{name}", redirect.New(log, storage))
	log.Info(fmt.Sprintf("start server on address:%s", cfg.Address))
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		os.Exit(1)
	}
	log.Info("server stopped")
}
func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler((os.Stdout), &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
