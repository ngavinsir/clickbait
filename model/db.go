package model

import (
	"database/sql"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
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

	return Connect(conn, "")
}

func Connect(conn string, path string) (*DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err := execSchema(db, path); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &DB{db}, nil
}

func execSchema(db *sql.DB, path string) error {
	if path == "" {
		path = "file://migrations"
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		path,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	m.Up()

	return nil
}
