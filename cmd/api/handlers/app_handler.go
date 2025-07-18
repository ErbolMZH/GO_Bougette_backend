package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) healthCheck(c echo.Context) error {
	healthCheckStruct := struct {
		health bool
	}{
		health: true,
	}
	return c.JSON(http.StatusOK, healthCheckStruct)
}
