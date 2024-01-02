package repository

import (
	database "github.com/edwincarlflores/mind/db"
	thoughtRepo "github.com/edwincarlflores/mind/repository/thought"
	userRepo "github.com/edwincarlflores/mind/repository/user"
)

type Repository struct {
	ThoughtRepository thoughtRepo.ThoughtRepository
	UserRepository    userRepo.UserRepository
}

func NewRepository(db database.DBInterface) *Repository {
	return &Repository{
		ThoughtRepository: thoughtRepo.NewThoughtDBRepository(db),
		UserRepository:    userRepo.NewUserDBRepository(db),
	}
}
