package flow

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Flow represents a focused work session configuration
type Flow struct {
	Name      string   `json:"name"`
	Resources []string `json:"resources"`
}

// Manager handles loading and saving flows
type Manager struct {
	Flows map[string]*Flow `json:"flows"`
	path  string
}

// NewManager creates a new flow manager
func NewManager() (*Manager, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, ".taskgo_flows.json")
	m := &Manager{
		Flows: make(map[string]*Flow),
		path:  path,
	}

	if err := m.Load(); err != nil {
		// If file doesn't exist, that's fine, we start empty
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return m, nil
}

// Load reads flows from disk
func (m *Manager) Load() error {
	data, err := os.ReadFile(m.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &m.Flows)
}

// Save writes flows to disk
func (m *Manager) Save() error {
	data, err := json.MarshalIndent(m.Flows, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.path, data, 0644)
}

// Create adds a new flow
func (m *Manager) Create(name string) error {
	if _, exists := m.Flows[name]; exists {
		return fmt.Errorf("flow '%s' already exists", name)
	}

	m.Flows[name] = &Flow{
		Name:      name,
		Resources: []string{},
	}

	return m.Save()
}

// AddResource adds a resource to a flow
func (m *Manager) AddResource(name string, resource string) error {
	flow, exists := m.Flows[name]
	if !exists {
		return fmt.Errorf("flow '%s' not found", name)
	}

	flow.Resources = append(flow.Resources, resource)
	return m.Save()
}

// Get returns a flow by name
func (m *Manager) Get(name string) (*Flow, error) {
	flow, exists := m.Flows[name]
	if !exists {
		return nil, fmt.Errorf("flow '%s' not found", name)
	}
	return flow, nil
}

// List returns all flow names
func (m *Manager) List() []string {
	keys := make([]string, 0, len(m.Flows))
	for k := range m.Flows {
		keys = append(keys, k)
	}
	return keys
}
