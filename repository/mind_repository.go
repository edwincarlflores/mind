package repository

import (
	"github.com/edwincarlflores/mind/entity"
)

type MindRepository interface {
	GetMindByUserID(userID int) (*entity.Mind, error)
}
