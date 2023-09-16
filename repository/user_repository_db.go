package repository

import (
	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/entity"
)

type UserDBRepository struct {
	DB database.DB
}

func NewUserDBRepository(db database.DB) *UserDBRepository {
	return &UserDBRepository{
		DB: db,
	}
}

func (r *UserDBRepository) GetUserByToken(token string) (*entity.User, error) {
	user := entity.User{}

	err := r.DB.Get(&user, "SELECT id, name, username, email, token FROM user WHERE token = ?", token)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserDBRepository) GetUserByUserName(userName string) (*entity.User, error) {
	user := entity.User{}

	err := r.DB.Get(&user, "SELECT id, name, username, email, token FROM user WHERE username = ?", userName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
