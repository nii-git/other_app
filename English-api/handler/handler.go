package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type Handler struct {
	logger slog.Logger
}

func NewHandler(logger slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Frequency API version β1.0")
	}
}

func (h *Handler) FrequencyHandlerFunc() echo.HandlerFunc {
	// todo:usecase
	return func(c echo.Context) error {
		h.logger.Info("frequency called")
		return c.String(http.StatusOK, c.QueryParam("test"))
	}
}
