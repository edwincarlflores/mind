package repository

import (
	"database/sql"
	"errors"
	"fmt"

	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/entity"
)

type ThoughtDBRepository struct {
	DB database.DBInterface
}

func NewThoughtDBRepository(db database.DBInterface) *ThoughtDBRepository {
	return &ThoughtDBRepository{
		DB: db,
	}
}

func (r *ThoughtDBRepository) GetAllThoughtsByUserID(userID string) ([]*entity.Thought, error) {
	thoughtsDB := []*dbThought{}

	err := r.DB.Select(&thoughtsDB, "SELECT id, body, description, kind, public FROM thought WHERE user_id = ?", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error when retrieving thoughts by user id %s: %w", userID, err)
	}

	thoughts := []*entity.Thought{}
	for _, t := range thoughtsDB {
		thoughts = append(thoughts, fromDBThought(t))
	}

	return thoughts, nil
}

func fromDBThought(thought *dbThought) *entity.Thought {
	return &entity.Thought{
		ID:          thought.ID,
		UUID:        thought.UUID,
		Body:        thought.Body,
		Description: thought.Description,
		Kind:        thought.Kind,
		Public:      thought.Public,
		UserID:      thought.UserID,
	}
}
