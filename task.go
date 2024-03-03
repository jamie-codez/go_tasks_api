package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type TaskService struct {
	store Store
}

func NewTaskService(s Store) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks", s.handleGetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", s.handleUpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", s.handlePatchTask).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", s.handleDeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks", s.handleDeleteTasks).Methods("DELETE")
}

func (s *TaskService) handleCreateTask(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handleGetTasks(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handleGetTask(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handleUpdateTask(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handlePatchTask(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handleDeleteTask(res http.ResponseWriter, req *http.Request) {

}

func (s *TaskService) handleDeleteTasks(res http.ResponseWriter, req *http.Request) {

}
