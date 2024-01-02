package repository

import "github.com/edwincarlflores/mind/entity"

type ThoughtRepository interface {
	GetAllThoughtsByUserID(userID string) ([]*entity.Thought, error)
}
