package mind

import (
	"errors"

	"github.com/edwincarlflores/mind/internal/core/domain/thought"
	"github.com/edwincarlflores/mind/internal/core/domain/user"
)

var ErrInvalidUser = errors.New("invalid user")

type Mind struct {
	User     *user.User
	Thoughts []*thought.Thought
}

func MindFromEntities(user *user.User, thoughts []*thought.Thought) (*Mind, error) {
	if user == nil {
		return nil, ErrInvalidUser
	}

	return &Mind{
		User:     user,
		Thoughts: thoughts,
	}, nil
}
