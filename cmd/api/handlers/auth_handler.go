package handlers

import (
	"bougette-backend/cmd/api/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) RegisterHandlers(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Bad reaquest")
	}
	c.Logger().Info(payload)
	return c.String(http.StatusBadRequest, "good request")
}
