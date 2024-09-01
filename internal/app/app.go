package app

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/hrvadl/book-service/internal/cfg"
	genredomain "github.com/hrvadl/book-service/internal/domain/genre"
	userdomain "github.com/hrvadl/book-service/internal/domain/user"
	"github.com/hrvadl/book-service/internal/storage/db"
	genrestorage "github.com/hrvadl/book-service/internal/storage/repo/genres"
	userstorage "github.com/hrvadl/book-service/internal/storage/repo/user"
	usertransport "github.com/hrvadl/book-service/internal/transport/http/user"
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
	db, err := db.NewSQL(a.cfg.PostgresDSN)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	a.log.Info("Successfully connected to the PostgreSQL")

	genresRepo := genrestorage.NewRepo(db)
	genresService := genredomain.NewService(genresRepo)

	usersRepo := userstorage.NewRepo(db)
	usersService := userdomain.NewService(usersRepo, genresService)
	usersHandler := usertransport.NewHandler(usersService)

	srv := fiber.New()
	srv.Use(logger.New())
	v1 := srv.Group("/v1")

	users := v1.Group("/users")
	users.Post("/", usersHandler.CreateUser)
	users.Get("/:id", usersHandler.GetByID)

	return srv.Listen(net.JoinHostPort(a.cfg.Host, a.cfg.Port))
}

func (a *App) GracefulStop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	signal := <-ch
	a.log.Info("Received stop signal. Terminating...", slog.Any("signal", signal))
	a.log.Info("Successfully terminated server. Bye!")
}
