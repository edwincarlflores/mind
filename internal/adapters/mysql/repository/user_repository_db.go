package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/edwincarlflores/mind/internal/adapters/mysql"
	"github.com/edwincarlflores/mind/internal/core/domain/user"
	"github.com/google/uuid"
)

type UserDBRepository struct {
	DB mysql.DBInterface
}

func NewUserDBRepository(db mysql.DBInterface) *UserDBRepository {
	return &UserDBRepository{
		DB: db,
	}
}

type dbUser struct {
	ID        uuid.UUID `db:"id"`
	UUID      uuid.UUID `db:"uuid"`
	Username  string    `db:"username"`
	AuthID    *string   `db:"auth_id"`
	AuthName  *string   `db:"auth_name"`
	Email     *string   `db:"email"`
	Name      string    `db:"name"`
	Provider  *string   `db:"provider"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
}

func (r *UserDBRepository) GetUserByUserName(userName string) (*user.User, error) {
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

func fromDBUser(dbu *dbUser) *user.User {
	return &user.User{
		ID:       dbu.ID,
		UUID:     dbu.UUID,
		Username: dbu.Username,
		AuthID:   dbu.AuthID,
		AuthName: dbu.AuthName,
		Email:    dbu.Email,
		Name:     dbu.Name,
		Provider: dbu.Provider,
	}
}
