package thought

import "github.com/google/uuid"

type Thought struct {
	ID          uuid.UUID
	UUID        uuid.UUID
	Body        string
	Description string
	Kind        string
	Public      bool
	UserID      uuid.UUID
	CreatedAt   string
	UpdatedAt   string
}
