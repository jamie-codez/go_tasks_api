package main

// Mocks
type MockStore struct{}

func (m *MockStore) CreateUser(user *User) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) CreateProject(project *Project) (*Project, error) {
	return &Project{}, nil
}

func (m *MockStore) CreateTask(task *Task) (*Task, error) {
	return &Task{}, nil
}
