package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDatabase(cfg *Config) (*sql.DB, func()) {
	conn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Error opening database " + err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Error pinging database " + err.Error())
	}

	return db, func() { db.Close() }
}
