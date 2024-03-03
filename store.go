package main

import "database/sql"

type Store interface {
	// User Methods
	// CreateUser creates a new user in the store
	CreateUser(task *User) (*User, error)
	// GetUsers returns all users from the store with pagination and filtering
	GetUsers(page int, size int) ([]*User, error)
	// GetUser returns a user from the store
	GetUser(id int) (*User, error)
	// UpdateUser updates a user in the store
	UpdateUser(user *User) (*User, error)
	// DeleteUser deletes a user from the store
	DeleteUser(id int) error

	// Project Methods
	// CreateProject creates a new project in the store
	CreateProject(project *Project) (*Project, error)
	// GetProjects returns all projects from the store with pagination and filtering
	GetProjects(page int, size int) ([]*Project, error)
	// GetProject returns a project from the store
	GetProject(id int) (*Project, error)
	// UpdateProject updates a project in the store
	UpdateProject(project *Project) (*Project, error)
	// DeleteProject deletes a project from the store
	DeleteProject(id int) error

	// Tasks Methods
	// CreateTask creates a new task in the store
	CreateTask(task *Task) (*Task, error)
	// GetTasks returns all tasks from the store with pagination and filtering
	GetTasks(page int, size int) ([]*Task, error)
	// GetTask returns a task from the store
	GetTask(id int) (*Task, error)
	// UpdateTask updates a task in the store
	UpdateTask(task *Task) (*Task, error)
	// DeleteTask deletes a task from the store
	DeleteTask(id int) error
	// Delete tasks by parsed id
	DeleteTasks(ids []int) error
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// User Methods
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
	s.db.Close()
	return user, nil
}

func (s *Storage) GetUsers(page int, size int) ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users LIMIT $1 OFFSET $2", size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*User{}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.EmailAddress, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	s.db.Close()
	return users, nil
}

func (s *Storage) GetUser(id int) (*User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.EmailAddress, &user.Password)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return user, nil
}

func (s *Storage) UpdateUser(user *User) (*User, error) {
	_, err := s.db.Exec("UPDATE users SET username = $1, first_name = $2, last_name = $3, email_address = $4, password = $5 WHERE id = $6", user.Username, user.FirstName, user.LastName, user.EmailAddress, user.Password, user.ID)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return user, nil
}

func (s *Storage) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	s.db.Close()
	return nil
}

// Project Methods
func (s *Storage) CreateProject(project *Project) (*Project, error) {
	rows, err := s.db.Exec("INSERT INTO projects (name,description,created_by) VALUES ($1, $2, $3) RETURNING id", project.Name, project.Description, project.CreatedBy)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	project.ID = int(id)
	s.db.Close()
	return project, nil
}

func (s *Storage) GetProjects(page int, size int) ([]*Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects LIMIT $1 OFFSET $2", size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := []*Project{}
	for rows.Next() {
		project := &Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedBy)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	s.db.Close()
	return projects, nil
}

func (s *Storage) GetProject(id int) (*Project, error) {
	row := s.db.QueryRow("SELECT * FROM projects WHERE id = $1", id)
	project := &Project{}
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedBy)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return project, nil
}

func (s *Storage) UpdateProject(project *Project) (*Project, error) {
	_, err := s.db.Exec("UPDATE projects SET name = $1, description = $2, created_by = $3 WHERE id = $4", project.Name, project.Description, project.CreatedBy, project.ID)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return project, nil
}

func (s *Storage) DeleteProject(id int) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}
	s.db.Close()
	return nil
}

// Tasks Methods
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
	s.db.Close()
	return task, nil
}

func (s *Storage) GetTasks(page int, size int) ([]*Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks LIMIT $1 OFFSET $2", size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []*Task{}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.ProjectID, &task.AssignedTo, &task.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	s.db.Close()
	return tasks, nil
}

func (s *Storage) GetTask(id int) (*Task, error) {
	row := s.db.QueryRow("SELECT * FROM tasks WHERE id = $1 RIGHT JOIN projects ON tasks.project_id = projects.id", id)
	task := &Task{}
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.Deadline)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return task, nil
}

func (s *Storage) UpdateTask(task *Task) (*Task, error) {
	_, err := s.db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, project_id = $4, assigned_user_id = $5, deadline = $6 WHERE id = $7", task.Title, task.Description, task.Status, task.ProjectID, task.AssignedTo, task.Deadline, task.ID)
	if err != nil {
		return nil, err
	}
	s.db.Close()
	return task, nil
}

func (s *Storage) DeleteTask(id int) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	s.db.Close()
	return nil
}

func (s *Storage) DeleteTasks(ids []int) error {
	for _, id := range ids {
		_, err := s.db.Exec("DELETE FROM tasks WHERE id = $1", id)
		if err != nil {
			return err
		}
	}
	s.db.Close()
	return nil
}
