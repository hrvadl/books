package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/hrvadl/book-service/internal/cfg"
	genredomain "github.com/hrvadl/book-service/internal/domain/genre"
	historydomain "github.com/hrvadl/book-service/internal/domain/history"
	userdomain "github.com/hrvadl/book-service/internal/domain/user"
	"github.com/hrvadl/book-service/internal/storage/db"
	genrestorage "github.com/hrvadl/book-service/internal/storage/repo/genres"
	historystorage "github.com/hrvadl/book-service/internal/storage/repo/history"
	userstorage "github.com/hrvadl/book-service/internal/storage/repo/user"
	"github.com/hrvadl/book-service/internal/transport/http/history"
	usertransport "github.com/hrvadl/book-service/internal/transport/http/user"
	"github.com/hrvadl/book-service/internal/transport/pubsub/subscribers/review"
)

const (
	timeout   = 5 * time.Second
	subName   = "user_added"
	topicName = subName
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
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pg, err := db.NewSQL(ctx, a.cfg.PostgresDSN)
	if err != nil {
		return fmt.Errorf("failed to init postgres: %w", err)
	}
	a.log.Info("Successfully connected to the PostgreSQL")

	firestore, err := db.NewFirestore(ctx, a.cfg.GCPKeypath, a.cfg.GCPProjectID, a.cfg.FirestoreDB)
	if err != nil {
		return fmt.Errorf("failed to init postgres: %w", err)
	}
	a.log.Info("Successfully connected to the Firestore")

	sub, err := review.NewSubscriber(ctx, review.SubOptions{
		Filename:         a.cfg.GCPKeypath,
		ProjectID:        a.cfg.GCPProjectID,
		SubscriptionName: subName,
	})
	if err != nil {
		return fmt.Errorf("failed to create pub/sub subcription: %w", err)
	}
	go sub.Subscribe(context.Background())
	a.log.Info("Successfully connected to the Pub/Sub")

	pub, err := review.NewPublisher(ctx, review.PubOptions{
		Filename:  a.cfg.GCPKeypath,
		ProjectID: a.cfg.GCPProjectID,
		Topic:     topicName,
	})
	if err != nil {
		return fmt.Errorf("failed to create pub/sub publisher: %w", err)
	}

	historyRepo := historystorage.NewRepo(firestore)
	historyService := historydomain.NewService(historyRepo)
	historyHandler := history.NewHandler(historyService)

	genresRepo := genrestorage.NewRepo(pg)
	genresService := genredomain.NewService(genresRepo)

	usersRepo := userstorage.NewRepo(pg)
	usersService := userdomain.NewService(usersRepo, genresService)
	usersHandler := usertransport.NewHandler(usersService, pub)

	srv := fiber.New()
	srv.Use(logger.New())
	v1 := srv.Group("/v1")

	users := v1.Group("/users")
	users.Post("/", usersHandler.CreateUser)
	users.Get("/:id", usersHandler.GetByID)

	histories := v1.Group("/history")
	histories.Post("/:userID/:bookID", historyHandler.Add)

	return srv.Listen(net.JoinHostPort(a.cfg.Host, a.cfg.Port))
}

func (a *App) GracefulStop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	signal := <-ch
	a.log.Info("Received stop signal. Terminating...", slog.Any("signal", signal))
	a.log.Info("Successfully terminated server. Bye!")
}
