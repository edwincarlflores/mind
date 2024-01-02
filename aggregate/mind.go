package aggregate

import (
	"errors"

	"github.com/edwincarlflores/mind/entity"
)

var ErrInvalidUser = errors.New("invalid user")

type Mind struct {
	User     *entity.User
	Thoughts []*entity.Thought
}

func MindFromEntities(user *entity.User, thoughts []*entity.Thought) (*Mind, error) {
	if user == nil {
		return nil, ErrInvalidUser
	}

	return &Mind{
		User:     user,
		Thoughts: thoughts,
	}, nil
}
