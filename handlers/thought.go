package handlers

import (
	"net/http"
	"strconv"

	"github.com/edwincarlflores/mind/db/repository"
	common "github.com/edwincarlflores/mind/templates/common"
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
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	thoughts, err := h.Repo.GetAllThoughts(mindID)
	if err != nil {
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	return utils.HTML(c, http.StatusOK, mindpage.Thoughts(thoughts))
}
