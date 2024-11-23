package users

import (
	"database/sql"
)

type Repository interface {
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetUser(id string) (User, error) {
	return User{}, nil
}

func (r repository) GetUsers() ([]User, error) {
	return []User{}, nil
}
