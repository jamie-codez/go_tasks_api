package main

import "database/sql"

type Store interface {
	// CreateUser creates a new user in the store
	CreateUser(task *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser(user *User) (*User, error) {
	rows, err := s.db.Exec("INSERT INTO users (username,first_name,last_name,email_address, password) VALUES ($1, $2) RETURNING id", user.Username, user.FirstName, user.LastName, user.EmailAddress, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (s *Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (title, description, status,project_id,assigned_user_id,deadline) VALUES ($1, $2, $3,$4,$5,$6) RETURNING id", task.Title, task.Description, task.Status, task.ProjectID, task.AssignedTo, task.Deadline)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = int(id)
	return task, nil
}
