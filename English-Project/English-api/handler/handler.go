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
	logger    slog.Logger
	usecase   usecase.Usecase
	validator model.Validator
}

func NewHandler(logger slog.Logger, usecase usecase.Usecase, validator model.Validator) *Handler {
	return &Handler{
		logger:    logger,
		usecase:   usecase,
		validator: validator,
	}
}

func (h *Handler) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Frequency API version β1.0")
	}
}

func (h *Handler) FrequencyHandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		h.logger.Debug("Frequency handler called")

		// TODO: validation
		limitparam := 100
		if c.QueryParam("Limit") != "" {
			limitparam, err = strconv.Atoi(c.QueryParam("Limit"))
			if err != nil {
				h.logger.Error("FrequencyHandlerFunc Invalid limit param: " + err.Error())
				return c.JSON(http.StatusBadRequest, model.Frequency_response{Error: "VALIDATION_ERROR", Body: nil})
			}
		}

		pageparam := 0
		if c.QueryParam("Page") != "" {
			pageparam, err = strconv.Atoi(c.QueryParam("Page"))
			if err != nil {
				h.logger.Error("FrequencyHandlerFunc Invalid pageparam: " + err.Error())
				return c.JSON(http.StatusBadRequest, model.Frequency_response{Error: "VALIDATION_ERROR", Body: nil})
			}
		}

		request := model.Frequency_request{
			Date:     c.QueryParam("Date"),
			Order:    c.QueryParam("Order"),
			Limit:    limitparam,
			Page:     pageparam,
			Provider: c.QueryParam("Provider"),
		}

		if err = h.validator.Validator.Struct(request); err != nil {
			h.logger.Error("FrequencyHandlerFunc ValidationError: " + err.Error())
			return c.JSON(http.StatusBadRequest, &model.Frequency_response{Error: "VALIDATION_ERROR", Body: nil})
		}

		res, err := h.usecase.FrequencyUsecase(request)

		if err != nil {
			h.logger.Error("FrequencyHandlerFunc UsecaseError:" + err.Error())
			return c.JSON(http.StatusNotFound, &model.Frequency_response{Error: "USECASE_ERROR", Body: nil})
		}
		return c.JSON(http.StatusOK, res)
	}
}
