package repository

import (
	"github.com/edwincarlflores/mind/internal/core/domain/thought"
	"github.com/edwincarlflores/mind/internal/core/domain/user"
)

type Repository struct {
	ThoughtRepository thought.ThoughtRepository
	UserRepository    user.UserRepository
}
