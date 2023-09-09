package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CreateDBConnection() *sqlx.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	db := sqlx.MustConnect("mysql", dbUrl)

	err := db.Ping()
	if err != nil {
		panic(err)
	} else {
		println("Connected to DB!")
	}

	return db
}

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
