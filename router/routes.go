package router

import (
	"github.com/edwincarlflores/mind/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HandleRender)
	e.GET("/thoughts", handlers.HandleGetAllThoughts)
}
