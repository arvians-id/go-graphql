package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgreSQL() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=root dbname=go_graphql sslmode=disable")
	if err != nil {
		return nil, err
	}

	db, err = databasePooling(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePooling(db *sql.DB) (*sql.DB, error) {
	db.SetMaxIdleConns(0)                                    // minimal connection
	db.SetMaxOpenConns(60)                                   // maximal connection
	db.SetConnMaxLifetime(time.Duration(3600) * time.Second) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(1000) * time.Second) // connection that can be used

	return db, nil
}
