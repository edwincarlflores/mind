package routes

import (
	"github.com/edwincarlflores/mind/internal/adapters/web/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, handler *handlers.Handler) {
	e.GET("/:username", handler.HandleGetMind)
}
