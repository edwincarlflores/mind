package service

import "github.com/edwincarlflores/mind/repository"

type Service struct {
	RenderService RenderServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		RenderService: NewRenderService(repo),
	}
}
