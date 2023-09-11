package repository

import (
	"github.com/jmoiron/sqlx"
)

type MindRepository struct {
	Conn *sqlx.DB
}

type Mind struct {
	ID         int    `db:"id"`
	OwnerToken string `db:"owner_token"`
}

func NewMindRepository(conn *sqlx.DB) *MindRepository {
	return &MindRepository{
		Conn: conn,
	}
}

func (r *MindRepository) GetMind(id int) (*Mind, error) {
	mindRows := []*Mind{}

	err := r.Conn.Select(&mindRows, "SELECT id, owner_token FROM mind WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if len(mindRows) == 0 {
		return nil, nil
	}

	return mindRows[0], nil
}
