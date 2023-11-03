package handler

import (
	"english-frequency/model"
	"english-frequency/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type Handler struct {
	logger  slog.Logger
	usecase usecase.Usecase
}

func NewHandler(logger slog.Logger, usecase usecase.Usecase) *Handler {
	return &Handler{
		logger:  logger,
		usecase: usecase,
	}
}

func (h *Handler) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Frequency API version β1.0")
	}
}

func (h *Handler) FrequencyHandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		h.logger.Debug("frequency handler called")

		// TODO: validation
		limitparam := 100
		if c.QueryParam("Limit") != "" {
			limitparam, err = strconv.Atoi(c.QueryParam("Limit"))
			if err != nil {
				h.logger.Error("FrequencyHandlerFunc Invalid limitparam: " + err.Error())
				return c.JSON(http.StatusBadRequest, model.Frequency_response{Error: err.Error(), Body: nil})
			}
		}

		pageparam := -1
		if c.QueryParam("Page") != "" {
			pageparam, err = strconv.Atoi(c.QueryParam("Page"))
			if err != nil {
				h.logger.Error("FrequencyHandlerFunc Invalid pageparam: " + err.Error())
				return c.JSON(http.StatusBadRequest, model.Frequency_response{Error: err.Error(), Body: nil})
			}
		}

		request := model.Frequency_request{
			Date:  c.QueryParam("Date"),
			Order: c.QueryParam("Order"),
			Limit: limitparam,
			Page:  pageparam,
		}

		h.usecase.FrequencyUsecase(request)
		return c.String(http.StatusOK, c.QueryParam("Page"))
	}
}
