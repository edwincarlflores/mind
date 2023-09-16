package repository

import (
	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/entity"
)

type ThoughtDBRepository struct {
	DB database.DB
}

func NewThoughtDBRepository(db database.DB) *ThoughtDBRepository {
	return &ThoughtDBRepository{
		DB: db,
	}
}

func (r *ThoughtDBRepository) GetAllThoughtsByUserID(userID int) ([]*entity.Thought, error) {
	thoughts := []*entity.Thought{}

	err := r.DB.Select(&thoughts, "SELECT id, body, description, kind, public FROM thought WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return thoughts, nil
}
