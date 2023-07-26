package main

import (
	"fmt"
	"go-api-1/server/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting our application...")
	e := echo.New()
	configs.ConnectDB()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	type event_schema struct {
		Name        string `json:"name"`
		VenueID     string `json:"venue_id"`
		Date        string `json:"date"`
		TicketCount int    `json:"ticket_count"`
	}

	// e.POST("/events", func(c echo.Context) error {
	// 	event := new(event_schema)
	// 	if err := c.Bind(event); err != nil {
	// 		return c.String(http.StatusBadRequest, "bad request")
	// 	}

	// 	fmt.Println(event)

	// 	res, err := DB.Collection("events").InsertOne(models.Event{event})

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	return c.JSON(http.StatusCreated, res)
	// })

	e.GET("/events", func(c echo.Context) error {
		return echo.ErrNotImplemented
	})

	e.Logger.Fatal(e.Start(":8080"))
}
