package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/IvkinNikita/GoAuthorization/sso/sso/cmd/internal/config"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production" // Добавлено, если это пропущено в оригинальном коде
)

func main() {
	// TODO: initilization config object
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: initilization logger
	log := setupLogger(cfg.Env)
	log.Info("string application",
		slog.String("env", cfg.Env),
		slog.Any("env", cfg),
		slog.Int("port", cfg.GRPC.Port))
	log.Debug("debug message")
	log.Error("error mesage")
	log.Warn("warn message")
	// TODO: initilization apppp (app)

	// TODO: start gRPC-server app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
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
	}
	return log

}
