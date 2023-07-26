package main

import (
	"fmt"
	"go-api-1/server/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting our application...")
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		fmt.Println("ping received")
		return c.String(http.StatusOK, "pong")
	})

	routes.EventRoute(e)
	routes.VenueRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
