package services

import (
	"fmt"

	"github.com/edwincarlflores/mind/internal/core/domain/mind"
	ports "github.com/edwincarlflores/mind/internal/ports"
)

type RenderServiceInterface interface {
	GetMindByUserName(userName string) (*mind.Mind, error)
}

type RenderService struct {
	repo *ports.Repository
}

func NewRenderService(repo *ports.Repository) *RenderService {
	return &RenderService{
		repo: repo,
	}
}

func (s *RenderService) GetMindByUserName(userName string) (*mind.Mind, error) {
	user, err := s.repo.UserRepository.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("invalid user")
	}

	thoughts, err := s.repo.ThoughtRepository.GetAllThoughtsByUserID(user.ID.String())
	if err != nil {
		return nil, err
	}

	if user == nil || thoughts == nil {
		return nil, fmt.Errorf("invalid user or thoughts")
	}

	mind, err := mind.MindFromEntities(user, thoughts)
	if err != nil {
		return nil, err
	}

	return mind, nil
}
