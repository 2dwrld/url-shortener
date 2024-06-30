package main

import (
	"fmt"
	"github.com/2dwrld/url-shortener/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config -> cleanenv
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)

	// TODO: init logger -> slog
	log := setupLogger(cfg.Env)
	// Благодаря функции With к каждому сообщению будет добавлено поле env с информацией о текущем окружении.
	// log = log.With(slog.String("env", cfg.Env))
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
	// TODO: init storage -> SQLite
	// TODO: init router - > "net/http"
	// TODO: run server
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
