package repository

import "github.com/jmoiron/sqlx"

type ThoughtRepository struct {
	Conn *sqlx.DB
}

type Thought struct {
	Body        string `db:"body"`
	Description string `db:"description"`
	Kind        string `db:"kind"`
	Public      bool   `db:"public"`
}

func NewThoughtRepository(conn *sqlx.DB) *ThoughtRepository {
	return &ThoughtRepository{
		Conn: conn,
	}
}

func (r *ThoughtRepository) GetAllThoughts(mindId int) ([]*Thought, error) {
	thoughts := []*Thought{}

	err := r.Conn.Select(&thoughts, "SELECT body, description, kind, public FROM thought WHERE mind_id = ?", mindId)
	if err != nil {
		return nil, err
	}

	return thoughts, nil
}
