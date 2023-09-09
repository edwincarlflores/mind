package router

import (
	"github.com/edwincarlflores/mind/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, thoughtHandler *handlers.ThoughtHandler, mindHandler *handlers.MindHandler) {
	e.GET("/:id", mindHandler.HandleGetMind)
	e.GET("/:id/thoughts", thoughtHandler.HandleGetAllThoughts)
}
