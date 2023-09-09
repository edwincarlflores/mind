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
