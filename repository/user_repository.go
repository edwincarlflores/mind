package repository

import "github.com/edwincarlflores/mind/entity"

type UserRepository interface {
	GetUserByToken(token string) (*entity.User, error)
	GetUserByUserName(userName string) (*entity.User, error)
}
