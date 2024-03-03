package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}
	service := NewTaskService(ms)
	t.Run("Should return and error is title is empty", func(t *testing.T) {
		task := &Task{
			Title: "",
		}
		b, err := json.Marshal(task)
		if err != nil {
			t.Fatal("Failed to marshal task")
		}
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}
