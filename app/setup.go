package app

import (
	"github.com/edwincarlflores/mind/router"
	"github.com/labstack/echo/v4"
)

func SetupAndRunApp() error {
	// Instantiate app server
	e := echo.New()

	// Serve static files from /static directory
	e.Static("/static", "static")

	// Setup routes
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}
