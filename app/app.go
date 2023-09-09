package app

import (
	"github.com/a-h/templ"
	"github.com/edwincarlflores/mind/db"
	mindpage "github.com/edwincarlflores/mind/templates/mind"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func App() {
	e := echo.New()

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, mindpage.MindPage())
	})

	e.GET("/thoughts", func(c echo.Context) error {
		return HTML(c, mindpage.Thoughts(db.Thoughts))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
