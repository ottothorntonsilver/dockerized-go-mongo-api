package routes

import (
	"go-api-1/server/controllers"

	"github.com/labstack/echo/v4"
)

func VenueRoute(e *echo.Echo) {
	e.GET("/venues", controllers.GetVenues)
}
