package handlers

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/edwincarlflores/mind/db/repository"
	mindpage "github.com/edwincarlflores/mind/templates/mind"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

type ThoughtHandler struct {
	Repo *repository.ThoughtRepository
}

func NewThoughtHandler(repo *repository.ThoughtRepository) *ThoughtHandler {
	return &ThoughtHandler{
		Repo: repo,
	}
}

func (h *ThoughtHandler) HandleRender(c echo.Context) error {
	mindID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return HTML(c, mindpage.MindPage(mindID))
}

func (h *ThoughtHandler) HandleGetAllThoughts(c echo.Context) error {
	mindID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	thoughts, err := h.Repo.GetAllThoughts(mindID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return HTML(c, mindpage.Thoughts(thoughts))
}
