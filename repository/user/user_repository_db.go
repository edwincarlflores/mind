package repository

import (
	"database/sql"
	"errors"
	"fmt"

	database "github.com/edwincarlflores/mind/db"
	"github.com/edwincarlflores/mind/entity"
)

type UserDBRepository struct {
	DB database.DBInterface
}

func NewUserDBRepository(db database.DBInterface) *UserDBRepository {
	return &UserDBRepository{
		DB: db,
	}
}

func (r *UserDBRepository) GetUserByUserName(userName string) (*entity.User, error) {
	userDB := dbUser{}

	err := r.DB.Get(&userDB, "SELECT id, name, username FROM user WHERE username = ?", userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving user %s from DB: %w", userName, err)
	}

	user := fromDBUser(&userDB)
	if user == nil {
		return nil, fmt.Errorf("invalid user conversion")
	}

	return user, nil
}

func fromDBUser(user *dbUser) *entity.User {
	return &entity.User{
		ID:       user.ID,
		UUID:     user.UUID,
		Username: user.Username,
		AuthID:   user.AuthID,
		AuthName: user.AuthName,
		Email:    user.Email,
		Name:     user.Name,
		Provider: user.Provider,
	}
}
