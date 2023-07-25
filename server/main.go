package main

import (
	"fmt"
	"net/http"

	"go-api-1/server/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting our application...")
	e := echo.New()
	configs.ConnectDB()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/events", func(c echo.Context) error {
		return echo.ErrNotImplemented
	})

	e.GET("/events", func(c echo.Context) error {
		return echo.ErrNotImplemented
	})

	e.Logger.Fatal(e.Start(":8080"))
}
