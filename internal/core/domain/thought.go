package domain

import "github.com/google/uuid"

type Thought struct {
	ID          uuid.UUID
	UUID        uuid.UUID
	Body        string
	Description string
	Kind        string
	Public      bool
	UserID      uuid.UUID
	CreatedAt   string
	UpdatedAt   string
}

type ThoughtRepository interface {
	GetAllThoughtsByUserID(userID string) ([]*Thought, error)
}
