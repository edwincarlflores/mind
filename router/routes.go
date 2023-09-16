package router

import (
	"github.com/edwincarlflores/mind/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, handler *handlers.Handler) {
	e.GET("/:username", handler.HandleGetMind)
	e.GET("/:username/thoughts", handler.HandleGetAllThoughts)
}
