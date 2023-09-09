package router

import (
	"github.com/edwincarlflores/mind/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, thoughtHandler *handlers.ThoughtHandler) {
	e.GET("/:id", thoughtHandler.HandleRender)
	e.GET("/:id/thoughts", thoughtHandler.HandleGetAllThoughts)
}
