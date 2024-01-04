package thought

type ThoughtRepository interface {
	GetAllThoughtsByUserID(userID string) ([]*Thought, error)
}
