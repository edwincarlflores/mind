package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	UUID     uuid.UUID
	Username string
	AuthID   *string
	AuthName *string
	Email    *string
	Name     string
	Provider *string
}

type UserRepository interface {
	GetUserByUserName(userName string) (*User, error)
}
