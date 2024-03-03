package main

import "time"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ProjectID   int       `json:"project_id"`
	AssignedTo  int       `json:"assigned_to"`
	Deadline    time.Time `json:"deadline"`
}
