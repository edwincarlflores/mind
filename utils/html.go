package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, statusCode int, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(statusCode)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}
