package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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
	// Get page and limit from request
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse page", nil)
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse size", nil)
		return
	}
	tasks, err := s.store.GetTasks(pageInt, sizeInt)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to get tasks", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Tasks retrieved successfully", tasks)
}

func (s *TaskService) handleGetTask(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse id", nil)
		return
	}
	task, err := s.store.GetTask(id)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to get task", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Task retrieved successfully", task)
}

func (s *TaskService) handleUpdateTask(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse id", nil)
		return
	}
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

	// Get task from DB
	task, err = s.store.GetTask(id)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to get task", nil)
		return
	}

	// Save to DB
	task, err = s.store.UpdateTask(task)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to update task", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Task updated successfully", task)
}

func (s *TaskService) handlePatchTask(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse id", nil)
		return
	}
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

	// Get task from DB
	task, err = s.store.GetTask(id)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to get task", nil)
		return
	}

	// Save to DB
	task, err = s.store.UpdateTask(task)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to update task", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Task updated successfully", task)
}

func (s *TaskService) handleDeleteTask(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse id", nil)
		return
	}
	err = s.store.DeleteTask(id)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to delete task", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Task deleted successfully", nil)
}

func (s *TaskService) handleDeleteTasks(res http.ResponseWriter, req *http.Request) {
	ids, ok := req.URL.Query()["ids"]
	if !ok || len(ids) < 1 {
		sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse ids", nil)
		return
	}
	var idList []int
	for _, id := range ids {
		i, err := strconv.Atoi(id)
		if err != nil {
			sendResponse(res, http.StatusBadRequest, "Bad Request", true, "Failed to parse id", nil)
			return
		}
		idList = append(idList, i)
	}
	err := s.store.DeleteTasks(idList)
	if err != nil {
		sendResponse(res, http.StatusInternalServerError, "Internal Server Error", true, "Failed to delete tasks", nil)
		return
	}
	sendResponse(res, http.StatusOK, "OK", false, "Tasks deleted successfully", nil)
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
