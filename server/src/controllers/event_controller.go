package controllers

import (
	"context"
	"go-api-1/server/configs"
	"go-api-1/server/models"
	"go-api-1/server/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventCollection *mongo.Collection = configs.GetCollection(configs.DB, "events")

func CreateEvent(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	event := new(models.EventSchema)
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

	res, err := eventCollection.InsertOne(ctx, event_model)

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

}
