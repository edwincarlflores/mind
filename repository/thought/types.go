package repository

import "github.com/google/uuid"

type dbThought struct {
	ID          uuid.UUID `db:"id"`
	UUID        uuid.UUID `db:"uuid"`
	Body        string    `db:"body"`
	Description string    `db:"description"`
	Kind        string    `db:"kind"`
	Public      bool      `db:"public"`
	UserID      uuid.UUID `db:"user_id"`
}
