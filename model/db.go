package model

import "database/sql"

// DB abstracts db connection instance.
type DB struct {
	*sql.DB
}

// InitDB opens new db connection by url.
func InitDB(url string) (*DB, error) {
    db, err := sql.Open("postgres", url)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}