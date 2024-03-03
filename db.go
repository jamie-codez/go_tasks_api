package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

type Config struct {
	host     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewPostgresStorage(config Config) *PostgresStorage {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", config.host, config.user, config.password, config.dbname, config.sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres DB")
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) Init() (*sql.DB, error) {
	// Create table if not exists
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}
	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}
	if err := s.createTasksTable(); err != nil {
		return nil, err
	}
	return s.db, nil
}

func (s *PostgresStorage) createUsersTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    username VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (email_address)
  );`)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (s *PostgresStorage) createProjectsTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS projects (
	id SERIAL NOT NULL,
	name VARCHAR(100) NOT NULL,
	user_id INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
  );`)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (s *PostgresStorage) createTasksTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL NOT NULL,
	title VARCHAR(100) NOT NULL,
	description TEXT NOT NULL,
	status ENUM('TODO', 'IN_PROGRESS','TESTING', 'DONE') NOT NULL,
	project_id INTEGER NOT NULL,
	assigned_user_id INTEGER NOT NULL,
	deadline TIMESTAMP NOT NULL,
	created_by INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	FOREIGN KEY (project_id) REFERENCES projects(id),
	FOREIGN KEY (assigned_user_id) REFERENCES users(id),
	FOREIGN KEY (created_by) REFERENCES users(id)
  );`)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
