package repository

import (
	"fmt"

	database "github.com/edwincarlflores/mind/db"
	entity "github.com/edwincarlflores/mind/entity"
)

type MindDBRepository struct {
	DB database.DB
}

func NewMindDBRepository(db database.DB) *MindDBRepository {
	return &MindDBRepository{
		DB: db,
	}
}

func (r *MindDBRepository) GetMindByUserID(userID int) (*entity.Mind, error) {
	mind := entity.Mind{}

	err := r.DB.Get(&mind, "SELECT id, user_id FROM mind WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get mind with user ID %d from DB: %w", userID, err)
	}

	return &mind, nil
}
