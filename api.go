package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	address string
	store   Store
}

func NewAPIServer(address string, store Store) *APIServer {
	return &APIServer{address: address, store: store}
}

func (s *APIServer) Start() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// Registering services
	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(router)
	log.Println("API Server started at", s.address)

	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
