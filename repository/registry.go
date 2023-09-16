package repository

import database "github.com/edwincarlflores/mind/db"

type Repository struct {
	MindRepisitory    MindRepository
	ThoughtRepository ThoughtRepository
	UserRepository    UserRepository
}

func NewRepository(db database.DB) *Repository {
	return &Repository{
		MindRepisitory:    NewMindDBRepository(db),
		ThoughtRepository: NewThoughtDBRepository(db),
		UserRepository:    NewUserDBRepository(db),
	}
}
