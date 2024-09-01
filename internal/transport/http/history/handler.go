package history

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/hrvadl/book-service/internal/domain/history"
)

func NewHandler(hs HistoryService) *Handler {
	return &Handler{
		history: hs,
	}
}

type HistoryService interface {
	Add(ctx context.Context, h history.ReadingHistory) (string, error)
}

type Handler struct {
	history HistoryService
}

type AddHistoryResponse struct {
	ID string `json:"id"`
}

func (h *Handler) Add(ctx fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid userID")
	}

	bookID, err := strconv.Atoi(ctx.Params("bookID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid bookID")
	}

	newHistory := history.ReadingHistory{BookID: bookID, UserID: userID}
	id, err := h.history.Add(ctx.Context(), newHistory)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(AddHistoryResponse{id})
}
