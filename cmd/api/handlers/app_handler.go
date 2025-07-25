package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthCheck struct {
	Health bool `json:"Health"`
}

func (h *Handler) HealthCheck(c echo.Context) error {
	healthCheckStruct := healthCheck{
		Health: true,
	}
	return c.JSON(http.StatusOK, healthCheckStruct)
}
