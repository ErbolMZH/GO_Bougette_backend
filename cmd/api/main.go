package main

import (
	"bougette-backend/cmd/api/handlers"
	"bougette-backend/cmd/api/middlewares"
	"bougette-backend/common"

	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := common.NewMysql()

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
	e.Use(middleware.Logger())
	e.Use(middlewares.CustomMiddleware)
	//e.GET("/", h.HealthCheck)
	app.routes(e, h)
	application := app
	fmt.Println(application)
	port := os.Getenv("APP_PORT")
	app_address := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(app_address))

}

//v,_:= time.Now().UTC().MarshalText()v
//log.PrintF("log.org/%s.json", string(v))
//os.WriteFile(`log,org/`+string(v)+`json`,0644)
