package user

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/hrvadl/book-service/internal/domain/genre"
	"github.com/hrvadl/book-service/internal/domain/user"
)

func NewHandler(us UsersService) *Handler {
	return &Handler{
		users: us,
	}
}

type UsersService interface {
	Create(ctx context.Context, cmd user.CreateUserCmd) (int, error)
	GetByID(ctx context.Context, id int) (*user.User, error)
}

type Handler struct {
	users UsersService
}

type createUserRequest struct {
	Name            string   `json:"name"            validate:"required"`
	Email           string   `json:"email"           validate:"required"`
	PreferredGenres []string `json:"preferredGenres" validate:"required,min=1"`
}

type createUserResponse struct {
	ID int `json:"id"`
}

func (h *Handler) CreateUser(ctx fiber.Ctx) error {
	req := new(createUserRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if len(req.PreferredGenres) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Preferred genres can not be empty")
	}

	cmd := user.CreateUserCmd{
		Name:           req.Name,
		Email:          req.Email,
		FavoriteGenres: req.PreferredGenres,
	}

	id, err := h.users.Create(ctx.Context(), cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(createUserResponse{id})
}

type getByIDResponse struct {
	ID              int           `json:"id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Email           string        `json:"email,omitempty"`
	PreferredGenres []genre.Genre `json:"preferredGenres,omitempty"`
}

func (h *Handler) GetByID(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid userID")
	}

	if id == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid userID")
	}

	u, err := h.users.GetByID(ctx.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.JSON(getByIDResponse{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		PreferredGenres: u.PreferredGenres,
	})
}
