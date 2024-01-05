package services

import "github.com/edwincarlflores/mind/internal/ports"

type Service struct {
	RenderService RenderServiceInterface
}

func NewService(repo *ports.Repository) *Service {
	return &Service{
		RenderService: NewRenderService(repo),
	}
}
