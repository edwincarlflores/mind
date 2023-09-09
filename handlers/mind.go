package handlers

import (
	"net/http"
	"strconv"

	"github.com/edwincarlflores/mind/db/repository"
	templates "github.com/edwincarlflores/mind/templates/mind"
	"github.com/edwincarlflores/mind/utils"
	"github.com/labstack/echo/v4"
)

type MindHandler struct {
	Repo *repository.MindRepository
}

func NewMindHandler(repo *repository.MindRepository) *MindHandler {
	return &MindHandler{
		Repo: repo,
	}
}

func (h *MindHandler) HandleGetMind(c echo.Context) error {
	mindID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return utils.HTML(c, templates.MindPage(mindID))
}
