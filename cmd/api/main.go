package main

import (
	"bougette-backend/cmd/api/handlers"
	"bougette-backend/common"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo
	hanler handlers.Handler
}

func main() {
	e := echo.New()
	err := godotenv.Load()

	db, err := common.NewMysql()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!Erbol")
	})

	h := handlers.Handler{
		DB: db,
	}
	app := Application{
		logger: e.Logger,
		server: e,
		hanler: h,
	}
	application := app
	fmt.Println(application)
	port := os.Getenv("APP_PORT")
	app_address := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(app_address))

}
