package routes

import (
	"go-api-1/server/controllers"

	"github.com/labstack/echo/v4"
)

func EventRoute(e *echo.Echo) {
	e.POST("/events", controllers.CreateEvent)
}
