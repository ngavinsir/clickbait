package model

import (
	"database/sql"
	"os"
)

// DB abstracts db connection instance.
type DB struct {
	*sql.DB
}

// InitDB opens new db connection by url.
func InitDB() (*DB, error) {
    conn := "dbname=clickbait host=localhost user=postgres password=postgres"
	if envConn := os.Getenv("DATABASE_URL"); envConn != "" {
		conn = envConn
	}

    db, err := sql.Open("postgres", conn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}