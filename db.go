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
	return s.db, nil
}
