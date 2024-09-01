package app

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"

	"github.com/hrvadl/book-service/internal/cfg"
	"github.com/hrvadl/book-service/internal/storage/db"
)

func New(cfg *cfg.Config, log *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

type App struct {
	log *slog.Logger
	cfg *cfg.Config
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	_, err := db.NewSQL(a.cfg.PostgresDSN)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	a.log.Info("Successfully connected to the PostgreSQL")

	srv := fiber.New()
	return srv.Listen(net.JoinHostPort(a.cfg.Host, a.cfg.Port))
}

func (a *App) GracefulStop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	signal := <-ch
	a.log.Info("Received stop signal. Terminating...", slog.Any("signal", signal))
	a.log.Info("Successfully terminated server. Bye!")
}
