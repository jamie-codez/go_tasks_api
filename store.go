package main

import "database/sql"

type Store interface {
	// CreateUser creates a new user in the store
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser() error {
	return nil
}
