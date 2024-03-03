package main

import "log"

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
	api := NewAPIServer(":8080", store)
	api.Start()
}
