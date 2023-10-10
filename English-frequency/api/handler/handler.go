package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Frequency API version β1.0")
	}
}

func FrequencyHandler(db *sql.DB) echo.HandlerFunc {
	// todo:usecase
	return func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("test"))
	}
}
