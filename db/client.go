package db

type Thought struct {
	ID          int
	Body        string
	Description string
	Kind        string
	Public      bool
}

var Thoughts = []Thought{
	{
		ID:          1,
		Body:        "Initial thought",
		Description: "Test tought",
		Public:      false,
	},
	{
		ID:          2,
		Body:        "Second thought",
		Description: "Test tought",
		Public:      true,
	},
	{
		ID:          3,
		Body:        "Third thought",
		Description: "Test tought",
		Public:      true,
	},
	{
		ID:          4,
		Body:        "Fourth thought",
		Description: "Test tought",
		Public:      false,
	},
}
