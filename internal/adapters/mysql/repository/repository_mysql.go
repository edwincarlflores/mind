package mysql

import (
	"github.com/edwincarlflores/mind/internal/adapters/mysql"
	"github.com/edwincarlflores/mind/internal/ports"
)

func NewDBRepository(db mysql.DBInterface) *ports.Repository {
	return &ports.Repository{
		ThoughtRepository: NewThoughtDBRepository(db),
		UserRepository:    NewUserDBRepository(db),
	}
}
