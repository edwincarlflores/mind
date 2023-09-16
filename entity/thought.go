package entity

type Thought struct {
	ID          string `db:"id"`
	Body        string `db:"body"`
	Description string `db:"description"`
	Kind        string `db:"kind"`
	Public      bool   `db:"public"`
}
