package domain

import (
	"errors"
)

var ErrInvalidUser = errors.New("invalid user")

type Mind struct {
	User     *User
	Thoughts []*Thought
}

func MindFromEntities(user *User, thoughts []*Thought) (*Mind, error) {
	if user == nil {
		return nil, ErrInvalidUser
	}

	return &Mind{
		User:     user,
		Thoughts: thoughts,
	}, nil
}
