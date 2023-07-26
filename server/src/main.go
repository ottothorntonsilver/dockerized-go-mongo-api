package main

import (
	"context"
	"fmt"
	"go-api-1/server/configs"
	"go-api-1/server/models"
	"go-api-1/server/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("Starting our application...")
	e := echo.New()
	DB := configs.ConnectDB()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	type event_schema struct {
		Name        string `json:"name"`
		VenueID     string `json:"venue_id"`
		Date        string `json:"date"`
		TicketCount int    `json:"ticket_count"`
	}

	e.POST("/events", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		event := new(event_schema)
		if err := c.Bind(&event); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		event_model := models.Event{
			Id:          primitive.NewObjectID(),
			Name:        event.Name,
			VenueID:     event.VenueID,
			Date:        event.Date,
			TicketCount: event.TicketCount,
		}

		res, err := configs.GetCollection(DB, "events").InsertOne(ctx, event_model)

		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				responses.DefaultResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &echo.Map{"data": err.Error()},
				},
			)

		}

		return c.JSON(http.StatusCreated, res)
	})

	e.GET("/events", func(c echo.Context) error {
		return echo.ErrNotImplemented
	})

	e.Logger.Fatal(e.Start(":8080"))
}
