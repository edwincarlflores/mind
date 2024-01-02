package handlers

import (
	"net/http"

	"github.com/edwincarlflores/mind/service"
	common "github.com/edwincarlflores/mind/templates/common"
	templates "github.com/edwincarlflores/mind/templates/mind"
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
		return render(c, http.StatusBadRequest, common.ErrorPage(err.Error()))
	}

	if mind.Thoughts == nil || mind.User == nil {
		return render(c, http.StatusNotFound, common.ErrorPage("This mind and it's memories doesn't exist"))
	}

	return render(c, http.StatusOK, templates.MindPage(mind))
}
