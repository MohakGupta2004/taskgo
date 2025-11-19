package task

import (
	"errors"
	"time"
)

// Storage interface to avoid circular dependency if we imported storage package here.
// Better design: define the interface where it is used.
type Storage interface {
	Load() ([]Task, error)
	Save(tasks []Task) error
}

type Manager struct {
	storage Storage
}

func NewManager(storage Storage) *Manager {
	return &Manager{storage: storage}
}

func (m *Manager) Add(title string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:        id,
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	return m.storage.Save(tasks)
}

func (m *Manager) List() ([]Task, error) {
	return m.storage.Load()
}

func (m *Manager) Update(id int, status TaskStatus) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			if status == StatusCompleted {
				now := time.Now()
				tasks[i].CompletedAt = &now
			} else {
				tasks[i].CompletedAt = nil
			}
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return m.storage.Save(tasks)
}

func (m *Manager) Remove(id int) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return errors.New("task not found")
	}

	return m.storage.Save(newTasks)
}
