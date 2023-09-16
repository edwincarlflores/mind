package entity

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Token    string `db:"token"`
}
