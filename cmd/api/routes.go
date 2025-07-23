package main

import (
	"bougette-backend/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func (app *Application) routes(e *echo.Echo, handler handlers.Handler) {

	e.GET("/", handler.HealthCheck)
}
