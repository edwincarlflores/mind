package service

import (
	"github.com/edwincarlflores/mind/entity"
	"github.com/edwincarlflores/mind/repository"
)

type RenderServiceInterface interface {
	GetMindByUserName(userName string) (*entity.Mind, error)
	GetAllThoughtsByUserName(userName string) ([]*entity.Thought, error)
}

type RenderService struct {
	repo *repository.Repository
}

func NewRenderService(repo *repository.Repository) *RenderService {
	return &RenderService{
		repo: repo,
	}
}

func (s *RenderService) GetMindByUserName(userName string) (*entity.Mind, error) {
	user, err := s.repo.UserRepository.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	mind, err := s.repo.MindRepisitory.GetMindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return mind, nil
}

func (s *RenderService) GetAllThoughtsByUserName(userName string) ([]*entity.Thought, error) {
	user, err := s.repo.UserRepository.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	thoughts, err := s.repo.ThoughtRepository.GetAllThoughtsByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return thoughts, nil
}
