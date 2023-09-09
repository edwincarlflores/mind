package app

import (
	"os"

	"github.com/edwincarlflores/mind/config"
	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/db/repository"
	"github.com/edwincarlflores/mind/handlers"
	"github.com/edwincarlflores/mind/router"
	"github.com/labstack/echo/v4"
)

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// Connect to database
	conn := database.CreateDBConnection()

	// Instantiate app server
	e := echo.New()

	// Serve static files from /static directory
	e.Static("/static", "static")

	// Inject db connection to thought repository
	thoughtRepo := repository.NewThoughtRepository(conn)

	// Inject thought repository to thought handler
	thoughtHandler := handlers.NewThoughtHandler(thoughtRepo)

	// Setup routes
	router.SetupRoutes(e, thoughtHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	return nil
}
