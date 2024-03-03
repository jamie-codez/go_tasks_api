package main

import (
	"fmt"
	"log"
)

func main() {
	config := Config{
		Envs.DBAddress,
		Envs.DBUser,
		Envs.DBPassword,
		Envs.DBName,
		"disable",
	}
	postgresStorage := NewPostgresStorage(config)
	db, err := postgresStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStorage(db)
	api := NewAPIServer(fmt.Sprintf("127.0.0.1:%s", Envs.Port), store)
	api.Start()
}
