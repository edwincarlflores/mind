package handlers

import (
	"net/http"

	"github.com/edwincarlflores/mind/service"
	common "github.com/edwincarlflores/mind/templates/common"
	templates "github.com/edwincarlflores/mind/templates/mind"
	"github.com/edwincarlflores/mind/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) HandleGetMind(c echo.Context) error {
	userName := c.Param("username")

	mind, err := h.svc.RenderService.GetMindByUserName(userName)
	if err != nil {
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	if mind == nil {
		return utils.HTML(c, http.StatusNotFound, common.ErrorPage("This mind and it's memories doesn't exist"))
	}

	return utils.HTML(c, http.StatusOK, templates.MindPage(userName))
}

func (h *Handler) HandleGetAllThoughts(c echo.Context) error {
	userName := c.Param("username")

	thoughts, err := h.svc.RenderService.GetAllThoughtsByUserName(userName)
	if err != nil {
		return utils.HTML(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	return utils.HTML(c, http.StatusOK, templates.Thoughts(thoughts))
}
