package ports

import (
	"github.com/edwincarlflores/mind/internal/core/domain"
)

type Repository struct {
	ThoughtRepository domain.ThoughtRepository
	UserRepository    domain.UserRepository
}
