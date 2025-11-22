package task

import (
	"errors"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/config"
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

func (m *Manager) Add(title string, group string, validity string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	if group == "" {
		group = "General"
	}

	// Calculate ValidUntil
	var validUntil *time.Time
	var duration time.Duration

	if validity != "" {
		d, err := time.ParseDuration(validity)
		if err == nil {
			duration = d
		}
	} else {
		// Load context for group validity
		ctx, err := config.LoadContext()
		if err == nil {
			if v, ok := ctx.GroupValidity[group]; ok {
				d, err := time.ParseDuration(v)
				if err == nil {
					duration = d
				}
			} else if group == "General" {
				// Default for General if not in config
				duration = 24 * time.Hour
			}
		}
	}

	if duration > 0 {
		t := time.Now().Add(duration)
		validUntil = &t
	}

	newTask := Task{
		ID:         id,
		Title:      title,
		Group:      group,
		Status:     StatusTodo,
		CreatedAt:  time.Now(),
		ValidUntil: validUntil,
	}

	tasks = append(tasks, newTask)
	return m.storage.Save(tasks)
}

func (m *Manager) CleanupExpired() error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	now := time.Now()
	changed := false

	for _, t := range tasks {
		if t.ValidUntil != nil && t.ValidUntil.Before(now) {
			changed = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if changed {
		return m.storage.Save(newTasks)
	}
	return nil
}

func (m *Manager) List() ([]Task, error) {
	if err := m.CleanupExpired(); err != nil {
		return nil, err
	}
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

func (m *Manager) UpdateTitle(id int, title string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = title
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

func (m *Manager) UpdateValidity(id int, validity string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			if validity == "" || validity == "none" {
				// Remove validity
				tasks[i].ValidUntil = nil
			} else {
				// Parse and set new validity
				d, err := time.ParseDuration(validity)
				if err != nil {
					return errors.New("invalid duration format")
				}
				newValidUntil := time.Now().Add(d)
				tasks[i].ValidUntil = &newValidUntil
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

func (m *Manager) UpdateGroupValidity(group string, validity string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	updated := false
	for i, t := range tasks {
		taskGroup := t.Group
		if taskGroup == "" {
			taskGroup = "General"
		}

		if taskGroup == group {
			if validity == "" || validity == "none" {
				// Remove validity
				tasks[i].ValidUntil = nil
			} else {
				// Parse and set new validity
				d, err := time.ParseDuration(validity)
				if err != nil {
					return errors.New("invalid duration format")
				}
				newValidUntil := time.Now().Add(d)
				tasks[i].ValidUntil = &newValidUntil
			}
			updated = true
		}
	}

	if !updated {
		// No tasks found in this group, but that's okay
		// The group validity will still be saved in context for future tasks
		return nil
	}

	return m.storage.Save(tasks)
}

func (m *Manager) RemoveByGroup(group string) error {
	tasks, err := m.storage.Load()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	for _, t := range tasks {
		taskGroup := t.Group
		if taskGroup == "" {
			taskGroup = "General"
		}
		if taskGroup != group {
			newTasks = append(newTasks, t)
		}
	}

	return m.storage.Save(newTasks)
}
