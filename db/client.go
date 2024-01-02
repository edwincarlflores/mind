package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
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

func CreateConnection() (*sqlx.DB, error) {
	dbName := os.Getenv("DB_NAME")
	dbUrl := os.Getenv("DATABASE_URL")
	driverName := os.Getenv("DRIVER_NAME")
	mainDbUrl := strings.Replace(dbUrl, "/"+dbName, "/", 1)

	mainDB, err := sqlx.Connect(driverName, mainDbUrl)
	if err != nil {
		return nil, fmt.Errorf("error establishing a connection to %s db: %w", driverName, err)
	}

	defer mainDB.Close()

	_, err = mainDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return nil, fmt.Errorf("error creating %s DB: %w", dbName, err)
	}

	db, err := sql.Open(driverName, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid DB URL: %w", err)
	}

	dbx := sqlx.NewDb(db, driverName)

	const maxRetries = 3
	for i := 0; i < maxRetries; i++ {
		err = dbx.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Duration((i+1)*100) * time.Millisecond)
	}
	if err != nil {
		return nil, fmt.Errorf("error establishing a connection to %s DB: %w", driverName, err)
	}

	return dbx, nil
}

type DB struct {
	*sqlx.DB
}

var _db *DB
var _dbSync sync.Once

func NewSqlxDB(db *sqlx.DB) *DB {
	return &DB{
		db,
	}
}

func Client() *DB {
	_dbSync.Do(func() {
		conn, err := CreateConnection()
		if err != nil {
			log.Fatalf("unable to initialize DB:S %w", err)
		}
		_db = NewSqlxDB(conn)
	})

	return _db
}
