package main

import (
	"log/slog"
	"os"

	"github.com/hrvadl/book-service/internal/app"
	"github.com/hrvadl/book-service/internal/cfg"
)

// TODO:
// 1. Add pub/sub publisher
func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg, err := cfg.NewFromEnv()
	if err != nil {
		log.Error("Failed to initialize config", slog.Any("err", err))
		os.Exit(1)
	}

	app := app.New(cfg, log)
	go app.MustRun()
	app.GracefulStop()
}
