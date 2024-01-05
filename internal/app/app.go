package app

import (
	"os"

	"github.com/edwincarlflores/mind/config"
	"github.com/edwincarlflores/mind/internal/adapters/mysql"
	repository "github.com/edwincarlflores/mind/internal/adapters/mysql/repository"
	"github.com/edwincarlflores/mind/internal/app/handlers"
	"github.com/edwincarlflores/mind/internal/app/routes"
	"github.com/edwincarlflores/mind/internal/core/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupAndRun() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// Connect to database
	conn := mysql.Client()

	// Instantiate app server
	e := echo.New()

	// Serve static files from /static directory
	e.Static("/static", "internal/app/static")

	e.Use(middleware.Logger())

	// Inject DB connection to repositories
	repo := repository.NewDBRepository(conn)

	// Inject repository to services
	service := services.NewService(repo)

	// Inject service to handlers
	handler := handlers.NewHandler(service)

	// Setup routes
	routes.SetupRoutes(e, handler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	return nil
}
