package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
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
	body, err := io.ReadAll(req.Body)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to read request body", nil)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to close request body", nil)
			return
		}
	}(req.Body)

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to unmarshal request body", nil)
		return
	}
	// Validation
	if err := validateTaskPayload(task); err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, err.Error(), nil)
		return
	}

	// Save to DB
	task, err = s.store.CreateTask(task)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to create task", nil)
		return
	}
	sendResponse(res, http.StatusCreated, "Created", false, "Task created successfully", task)
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

func validateTaskPayload(t *Task) error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}
	if t.Status == "" {
		return errors.New("status is required")
	}
	if t.ProjectID == 0 {
		return errors.New("ProjectID is required")
	}
	if t.AssignedTo == 0 {
		return errors.New("AssignedTo is required")
	}
	if t.Deadline.IsZero() {
		return errors.New("deadline is required")
	}
	return nil
}
