package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Print("we are in the custom middleware")
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
