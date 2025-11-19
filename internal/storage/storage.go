package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/MohakGupta2004/taskgo/internal/task"
)

type Storage interface {
	Load() ([]task.Task, error)
	Save(tasks []task.Task) error
}

type JSONStorage struct {
	FilePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{FilePath: filePath}
}

func (s *JSONStorage) Load() ([]task.Task, error) {
	if _, err := os.Stat(s.FilePath); os.IsNotExist(err) {
		return []task.Task{}, nil
	}

	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *JSONStorage) Save(tasks []task.Task) error {
	dir := filepath.Dir(s.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
