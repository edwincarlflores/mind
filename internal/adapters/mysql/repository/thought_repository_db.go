package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/edwincarlflores/mind/internal/adapters/mysql"
	"github.com/edwincarlflores/mind/internal/core/domain/thought"
	"github.com/google/uuid"
)

type ThoughtDBRepository struct {
	DB mysql.DBInterface
}

func NewThoughtDBRepository(db mysql.DBInterface) *ThoughtDBRepository {
	return &ThoughtDBRepository{
		DB: db,
	}
}

type dbThought struct {
	ID          uuid.UUID `db:"id"`
	UUID        uuid.UUID `db:"uuid"`
	Body        string    `db:"body"`
	Description string    `db:"description"`
	Kind        string    `db:"kind"`
	Public      bool      `db:"public"`
	UserID      uuid.UUID `db:"user_id"`
}

func (r *ThoughtDBRepository) GetAllThoughtsByUserID(userID string) ([]*thought.Thought, error) {
	thoughtsDB := []*dbThought{}

	err := r.DB.Select(&thoughtsDB, "SELECT id, body, description, kind, public FROM thought WHERE user_id = ?", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error when retrieving thoughts by user id %s: %w", userID, err)
	}

	thoughts := []*thought.Thought{}
	for _, t := range thoughtsDB {
		thoughts = append(thoughts, fromDBThought(t))
	}

	return thoughts, nil
}

func fromDBThought(dbt *dbThought) *thought.Thought {
	return &thought.Thought{
		ID:          dbt.ID,
		UUID:        dbt.UUID,
		Body:        dbt.Body,
		Description: dbt.Description,
		Kind:        dbt.Kind,
		Public:      dbt.Public,
		UserID:      dbt.UserID,
	}
}
