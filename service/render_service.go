package service

import (
	"fmt"

	"github.com/edwincarlflores/mind/aggregate"
	"github.com/edwincarlflores/mind/repository"
)

type RenderServiceInterface interface {
	GetMindByUserName(userName string) (*aggregate.Mind, error)
}

type RenderService struct {
	repo *repository.Repository
}

func NewRenderService(repo *repository.Repository) *RenderService {
	return &RenderService{
		repo: repo,
	}
}

func (s *RenderService) GetMindByUserName(userName string) (*aggregate.Mind, error) {
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

	mind, err := aggregate.MindFromEntities(user, thoughts)
	if err != nil {
		return nil, err
	}

	return mind, nil
}
