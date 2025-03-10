package main

import (
	"log/slog"
	"os"
	"restapi/internal/config"
	"restapi/internal/storage/postgresql"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.InitConfig()
	log := setUpLogger(cfg.Env)
	log.Info("start application", slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Addres))
	storage, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage")
		return
	}
	_ = storage
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
