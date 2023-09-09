package handlers

import (
	"net/http"
	"strconv"

	"github.com/edwincarlflores/mind/db/repository"
	mindpage "github.com/edwincarlflores/mind/templates/mind"
	"github.com/edwincarlflores/mind/utils"
	"github.com/labstack/echo/v4"
)

type ThoughtHandler struct {
	Repo *repository.ThoughtRepository
}

func NewThoughtHandler(repo *repository.ThoughtRepository) *ThoughtHandler {
	return &ThoughtHandler{
		Repo: repo,
	}
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

	return utils.HTML(c, mindpage.Thoughts(thoughts))
}
