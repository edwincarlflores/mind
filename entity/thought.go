package entity

import "github.com/google/uuid"

type Thought struct {
	ID          uuid.UUID `db:"id"`
	UUID        uuid.UUID `db:"uuid"`
	Body        string    `db:"body"`
	Description string    `db:"description"`
	Kind        string    `db:"kind"`
	Public      bool      `db:"public"`
	UserID      uuid.UUID `db:"user_id"`
	CreatedAt   string    `db:"created_at"`
	UpdatedAt   string    `db:"updated_at"`
}
