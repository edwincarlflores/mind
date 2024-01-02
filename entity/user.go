package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `db:"id"`
	UUID     uuid.UUID `db:"uuid"`
	Username string    `db:"username"`
	AuthID   *string   `db:"auth_id"`
	AuthName *string   `db:"auth_name"`
	Email    *string   `db:"email"`
	Name     string    `db:"name"`
	Provider *string   `db:"provider"`
}
