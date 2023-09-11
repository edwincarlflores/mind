package handlers

import (
	"net/http"
	"strconv"

	"github.com/edwincarlflores/mind/db/repository"
	common "github.com/edwincarlflores/mind/templates/common"
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
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	mind, err := h.Repo.GetMind(mindID)
	if err != nil {
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	if mind == nil {
		return utils.HTML(c, http.StatusNotFound, common.ErrorPage("This mind and it's memories doesn't exist"))
	}

	return utils.HTML(c, http.StatusOK, templates.MindPage(mindID))
}
