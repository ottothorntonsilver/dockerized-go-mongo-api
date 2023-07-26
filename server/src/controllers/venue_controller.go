package controllers

import (
	"context"
	"go-api-1/server/configs"
	"go-api-1/server/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var venueCollection *mongo.Collection = configs.GetCollection(configs.DB, "venues")

func GetVenues(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := eventCollection.Find(ctx, bson.D{})

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

	return c.JSON(http.StatusOK, res)

}
