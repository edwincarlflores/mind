package server

import (
	"os"

	"github.com/edwincarlflores/mind/config"
	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/handlers"
	"github.com/edwincarlflores/mind/repository"
	"github.com/edwincarlflores/mind/router"
	"github.com/edwincarlflores/mind/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupAndRunServer() error {
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

	e.Use(middleware.Logger())

	// Inject DB connection to repositories
	repo := repository.NewRepository(conn)

	// Inject repository to services
	service := service.NewService(repo)

	// Inject service to handlers
	handler := handlers.NewHandler(service)

	// Setup routes
	router.SetupRoutes(e, handler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	return nil
}
