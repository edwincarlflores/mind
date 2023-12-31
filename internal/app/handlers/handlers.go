package handlers

import (
	"net/http"

	views "github.com/edwincarlflores/mind/internal/app/views/mind"
	shared "github.com/edwincarlflores/mind/internal/app/views/shared"
	"github.com/edwincarlflores/mind/internal/core/services"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *services.Service
}

func NewHandler(svc *services.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) HandleGetMind(c echo.Context) error {
	userName := c.Param("username")

	mind, err := h.svc.RenderService.GetMindByUserName(userName)
	if err != nil {
		return render(c, http.StatusBadRequest, shared.ErrorPage(err.Error()))
	}

	if mind.Thoughts == nil || mind.User == nil {
		return render(c, http.StatusNotFound, shared.ErrorPage("This mind and it's memories doesn't exist"))
	}

	return render(c, http.StatusOK, views.MindPage(mind))
}
