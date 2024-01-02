package repository

import "github.com/edwincarlflores/mind/entity"

type UserRepository interface {
	GetUserByUserName(userName string) (*entity.User, error)
}
