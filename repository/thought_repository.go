package repository

import "github.com/edwincarlflores/mind/entity"

type ThoughtRepository interface {
	GetAllThoughtsByUserID(userID int) ([]*entity.Thought, error)
}
